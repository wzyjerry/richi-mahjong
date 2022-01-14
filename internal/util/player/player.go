// player 玩家
// 包括点棒、手牌、弃牌、吃碰牌、座次
package player

import (
	"github.com/wzyjerry/mahjong/data"
	"github.com/wzyjerry/mahjong/internal/ent/schema/common"
	"github.com/wzyjerry/mahjong/internal/util"
	"github.com/wzyjerry/mahjong/internal/util/deck"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

type (
	Player interface {
		setTenbo(tenbo int64) Player

		// 当前座次
		Seat() common.Fon
		// 轮庄
		SeatChange()
		// 当前点棒
		Tenbo() int64
		// 点棒转移，向`to`支付`tenbo`点棒
		TenboTransfer(to Player, tenbo uint64)
		// 抓牌，返回当前张
		Draw(pack *deck.Deck) *deck.Tile

		// 检查是否可以吃牌，返回候选集
		TryChi(last *deck.Tile) [][2]string
		// 检查是否可以碰牌，返回候选集
		TryPon(last *deck.Tile) [][2]string
		// 检查是否可以开杠，只检查打出牌是否可开杠（大明杠）
		TryMinKan(last *deck.Tile) bool

		// （领空）检查拔北，返回不同种北牌
		TryNuku() []string
		// （领空）检查开杠，返回候选集（每种一张）
		TryKan() []*deck.Tile
		// （领空）14张牌检查立直
		TryRichi() bool

		// 13张牌检查铳牌
		CheckTenpai() []*deck.Tile

		// 13张牌检查tile是否可以和牌
		TryAgari(tile *deck.Tile) bool

		// 报菜名
		GetScore(status *common.AgariStatus, wanpai *deck.Wanpai) (int64, int64, []*common.Yaku)

		// TODO: 食替判断

		String() string

		// TODO: 测试用
		SetRon(t *deck.Tile)
		Set()
	}
	player struct {
		seat    common.Fon
		tenbo   int64
		players int64

		// hand.Back()代表最后摸牌
		hand *deck.Deck
		// 手牌映射，方便进行吃碰检查
		handMapping map[int64][]*deck.Tile

		discard  *deck.Deck
		furoList []*deck.Ordered

		// 拔北数量
		nuku int64

		// 标记状态
		richi   bool
		wrichi  bool
		ippatsu bool
		ron     *deck.Tile
	}
)

func (p *player) Set() {
	p.richi = true
	p.ippatsu = true
	p.nuku = 2
}

func (p *player) String() string {
	builder := strings.Builder{}
	builder.WriteString(util.Fon2String(p.seat))
	builder.WriteString(": ")
	builder.WriteString(util.MPSZCompress(p.hand.ShortString()))
	return builder.String()
}

func New(seat common.Fon, tenbo int64, players int64) Player {
	return &player{
		seat:        seat,
		tenbo:       tenbo,
		players:     players,
		hand:        deck.NewDeck(players),
		handMapping: make(map[int64][]*deck.Tile),
		discard:     deck.NewDeck(players),
		furoList:    make([]*deck.Ordered, 0),
	}
}

// 七对子判断，相邻一致且间隔不一致
func testChitoitsu(sorted []*deck.Tile) bool {
	for i := 0; i < 7; i++ {
		if !sorted[i<<1].Equal(sorted[i<<1|1]) {
			return false
		}
		if i < 6 && sorted[i<<1].Equal(sorted[i<<2]) {
			return false
		}
	}
	return true
}

// 国士判断，包含全部19牌，返回重复枚
func testKokushi(sorted []*deck.Tile) (bool, *deck.Tile) {
	pos := 0
	var tile *deck.Tile
	for i := range sorted {
		if sorted[i].Equal(Kokushi[pos]) {
			pos++
			continue
		} else if i > 0 && sorted[i].Equal(sorted[i-1]) {
			tile = sorted[i]
			continue
		}
		break
	}
	return pos == 14, tile
}

// 拆分排序手牌到4种花色
func splitSorted(sorted []*deck.Tile) []int {
	results := make([]int, 0, 4)
	hand := make(map[common.Suit][]int)
	for suit := common.Suit_SUIT_MANZU; suit <= common.Suit_SUIT_TSUPAI; suit++ {
		hand[suit] = make([]int, 9)
	}
	for _, t := range sorted {
		hand[t.GetSuit()][t.GetNumber()-1]++
	}
	for suit := common.Suit_SUIT_MANZU; suit <= common.Suit_SUIT_TSUPAI; suit++ {
		result := 0
		for _, num := range hand[suit] {
			result = result*5 + num
		}
		results = append(results, result)
	}
	return results
}

func testNormal(sorted []*deck.Tile) bool {
	hasEx := false
	indicates := splitSorted(sorted)
	for i := 0; i < 3; i++ {
		if !data.Normal.ContainsInt(indicates[i]) {
			if data.NormalEx.ContainsInt(indicates[i]) && !hasEx {
				hasEx = true
			} else {
				return false
			}
		}
	}
	if !data.Tsupai.ContainsInt(indicates[3]) {
		if data.TsupaiEx.ContainsInt(indicates[3]) && !hasEx {
			hasEx = true
		} else {
			return false
		}
	}
	return hasEx
}

func (p *player) TryRichi() bool {
	hand := p.hand.Tiles()
	for i := 0; i < 14; i++ {
		tiles := make([]*deck.Tile, len(hand))
		copy(tiles, hand)
		tiles = append(tiles[:i], tiles[i+1:]...)
		if len(testTenpai(tiles)) > 0 {
			return true
		}
	}
	return false
}

func (p *player) TryAgari(tile *deck.Tile) bool {
	tiles := p.hand.Tiles()
	if len(tiles) == 14 {
		tiles = tiles[:13]
	}
	for _, t := range testTenpai(p.hand.Tiles()) {
		if t.Equal(tile) {
			return true
		}
	}
	return false
}

// 13张牌检查听牌，返回铳牌集合
func testTenpai(tiles []*deck.Tile) []*deck.Tile {
	candidates := make([]*deck.Tile, 0)
	for suit := common.Suit_SUIT_MANZU; suit <= common.Suit_SUIT_TSUPAI; suit++ {
		nMax := int64(9)
		if suit == common.Suit_SUIT_TSUPAI {
			nMax = 7
		}
		for number := int64(1); number <= nMax; number++ {
			tile := GenTile(suit, number)
			sortedEx := make([]*deck.Tile, len(tiles)+1)
			copy(sortedEx, tiles)
			sortedEx[len(tiles)] = tile
			sort.Slice(sortedEx, func(i, j int) bool {
				return sortedEx[i].Less(sortedEx[j])
			})
			if testNormal(sortedEx) {
				candidates = append(candidates, tile)
			} else if len(sortedEx) == 14 {
				if testChitoitsu(sortedEx) {
					candidates = append(candidates, tile)
				} else if ok, _ := testKokushi(sortedEx); ok {
					candidates = append(candidates, tile)
				}
			}
		}
	}
	return candidates
}

func (p *player) CheckTenpai() []*deck.Tile {
	return testTenpai(p.hand.Tiles())
}

func (p *player) TryKan() []*deck.Tile {
	candidates := make([]*deck.Tile, 0)
	// 分为两种，一种handMapping中有4张牌，直接开杠；一种副露列表中存在碰牌加杠
	// 4张牌
	for _, lst := range p.handMapping {
		if len(lst) == 4 {
			candidates = append(candidates, lst[0])
		}
	}
	// 杠牌列表存在
	for _, furo := range p.furoList {
		if furo.Kind() == common.OrderedKind_ORDERED_KIND_PON {
			t := furo.Front()
			if lst := p.handMapping[p.mappingKey(t.GetSuit(), t.GetNumber())]; len(lst) == 1 {
				candidates = append(candidates, lst[0])
			}
		}
	}
	return candidates
}

func genKind(t *deck.Tile) byte {
	var kind byte
	if t.GetChi() {
		kind |= 1
	}
	if t.GetTransparent() {
		kind |= 2
	}
	return kind
}

// 从手牌中取一张牌，如果存在，返回`每种`牌的一个名字，红宝牌、透明状态都会导致不同种牌出现……
func (p *player) getTile(suit common.Suit, number int64) []string {
	result := make([]string, 0)
	// 种类记录，第1位标明是否红宝牌，第2位标明是否透明
	state := make(map[byte]struct{})

	for _, tile := range p.handMapping[p.mappingKey(suit, number)] {
		kind := genKind(tile)
		if _, ok := state[kind]; !ok {
			state[kind] = struct{}{}
			result = append(result, tile.GetName())
		}
	}
	return result
}

func (p *player) TryNuku() []string {
	// 只有3麻可以拔北
	if p.players != 3 {
		return make([]string, 0)
	}
	return p.getTile(common.Suit_SUIT_TSUPAI, 4)
}

// 返回a与b的配对结果，a在前b在后
func (p *player) crossProduct(a, b []string) [][2]string {
	result := make([][2]string, 0, len(a)*len(b))
	for _, aa := range a {
		for _, bb := range b {
			result = append(result, [2]string{aa, bb})
		}
	}
	return result
}

func (p *player) TryChi(last *deck.Tile) [][2]string {
	candidates := make([][2]string, 0)
	// 三麻不允许吃
	if p.players == 3 {
		return candidates
	}
	// 字牌不会被吃
	if last.GetSuit() == common.Suit_SUIT_TSUPAI {
		return candidates
	}
	number := last.GetNumber()
	// xxO
	if 3 <= number {
		a := p.getTile(last.GetSuit(), number-2)
		b := p.getTile(last.GetSuit(), number-1)
		candidates = append(candidates, p.crossProduct(a, b)...)
	}
	// xOx
	if 2 <= number && number <= 8 {
		a := p.getTile(last.GetSuit(), number-1)
		b := p.getTile(last.GetSuit(), number+1)
		candidates = append(candidates, p.crossProduct(a, b)...)
	}
	// Oxx
	if number <= 7 {
		a := p.getTile(last.GetSuit(), number+1)
		b := p.getTile(last.GetSuit(), number+2)
		candidates = append(candidates, p.crossProduct(a, b)...)
	}
	return candidates
}

type namedTile struct {
	name string
	kind byte
}

func (p *player) TryPon(last *deck.Tile) [][2]string {
	candidates := make([][2]string, 0)
	key := p.mappingKey(last.GetSuit(), last.GetNumber())
	if len(p.handMapping[key]) < 2 {
		return candidates
	}
	tiles := make([]*namedTile, 0, len(p.handMapping[key]))
	for _, t := range p.handMapping[key] {
		tiles = append(tiles, &namedTile{
			name: t.GetName(),
			kind: genKind(t),
		})
	}
	sort.SliceStable(tiles, func(i, j int) bool {
		if tiles[i].kind == tiles[j].kind {
			return tiles[i].name < tiles[j].name
		}
		return tiles[i].kind < tiles[j].kind
	})
	used := make(map[byte]struct{})
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			key := tiles[i].kind<<2 | tiles[j].kind
			if _, ok := used[key]; !ok {
				used[key] = struct{}{}
				candidates = append(candidates, [2]string{tiles[i].name, tiles[j].name})
			}
		}
	}
	return candidates
}

func (p *player) TryMinKan(last *deck.Tile) bool {
	return len(p.handMapping[p.mappingKey(last.GetSuit(), last.GetNumber())]) == 3
}

func (p *player) Seat() common.Fon {
	return p.seat
}

func (p *player) SeatChange() {
	seat := int64(p.seat)%p.players + 1
	p.seat = common.Fon(seat)
}

func (p *player) Tenbo() int64 {
	return p.tenbo
}

func (p *player) setTenbo(tenbo int64) Player {
	p.tenbo = tenbo
	return p
}

func (p *player) TenboTransfer(to Player, tenbo uint64) {
	deltaTenbo := int64(tenbo)
	p.setTenbo(p.Tenbo() - deltaTenbo)
	to.setTenbo(to.Tenbo() + deltaTenbo)
}

func (p *player) mappingKey(suit common.Suit, number int64) int64 {
	return int64(suit)*10 + number
}

func (p *player) Draw(pack *deck.Deck) *deck.Tile {
	t := pack.Front()
	pack.Remove(t)
	p.hand.PushBack(t)
	key := p.mappingKey(t.GetSuit(), t.GetNumber())
	p.handMapping[key] = append(p.handMapping[key], t)
	return t
}

var Kokushi = genKokushi()

func genKokushi() []*deck.Tile {
	result := make([]*deck.Tile, 0, 13)
	result = append(result, GenTile(common.Suit_SUIT_MANZU, 1))
	result = append(result, GenTile(common.Suit_SUIT_MANZU, 9))
	result = append(result, GenTile(common.Suit_SUIT_PINZU, 1))
	result = append(result, GenTile(common.Suit_SUIT_PINZU, 9))
	result = append(result, GenTile(common.Suit_SUIT_SOUZU, 1))
	result = append(result, GenTile(common.Suit_SUIT_SOUZU, 9))
	for i := int64(1); i <= 7; i++ {
		result = append(result, GenTile(common.Suit_SUIT_TSUPAI, i))
	}
	return result
}

// 返回只包含花色和数字的单张牌
func GenTile(suit common.Suit, number int64) *deck.Tile {
	return &deck.Tile{
		Tile: &common.Tile{
			Name:        util.PString(strconv.Itoa(rand.Int())),
			Suit:        &suit,
			Number:      &number,
			Transparent: util.PBool(false),
			Chi:         util.PBool(false),
		},
	}
}

// TODO: 测试用
func (p *player) SetRon(t *deck.Tile) {
	p.ron = t
}
