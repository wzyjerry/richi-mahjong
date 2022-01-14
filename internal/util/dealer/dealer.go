// 发牌员决定对局的开始和结束
// 使用配置初始化
// 每轮开始从Manager处领取一副全新麻将
package dealer

import (
	"github.com/wzyjerry/mahjong/internal/ent/schema/common"
	"github.com/wzyjerry/mahjong/internal/util"
	"github.com/wzyjerry/mahjong/internal/util/config"
	"github.com/wzyjerry/mahjong/internal/util/deck"
	"github.com/wzyjerry/mahjong/internal/util/player"
	"strconv"
	"strings"
)

type Dealer struct {
	name   string
	config *config.Config

	players []player.Player
	ken     common.Fon // 当前场风
	round   int        // 第几局
	richibo int        // 立直棒
	tsumibo int        // 场棒

	pack   *deck.Deck   // 壁牌
	wanpai *deck.Wanpai // 王牌

	current int64      // 当前玩家
	last    *deck.Tile // 最后打出的牌

	shortState string // 记录配牌后的牌山
	md5        string // 记录对应MD5
}

// TODO: 测试使用
func (d *Dealer) Wanpai() *deck.Wanpai {
	return d.wanpai
}

func New(name string, config *config.Config) *Dealer {
	builder := strings.Builder{}
	builder.WriteString(name)
	builder.WriteString("·")
	switch *config.Players {
	case 3:
		builder.WriteString("三人")
	case 4:
		builder.WriteString("四人")
	}
	switch *config.Rounds {
	case 1:
		builder.WriteString("东")
	case 2:
		builder.WriteString("南")
	case 3:
		builder.WriteString("一局")
	}
	return &Dealer{
		name:    builder.String(),
		config:  config,
		players: make([]player.Player, 0, 4),
	}
}

func (d *Dealer) String() string {
	builder := strings.Builder{}
	builder.WriteString(d.Name())
	builder.WriteString("\n")
	builder.WriteString(d.RoundDesc())
	builder.WriteString("\n")
	builder.WriteString(d.shortState)
	builder.WriteString("\n")
	builder.WriteString(d.md5)
	builder.WriteString("\n")
	builder.WriteString(d.pack.String())
	builder.WriteString("\n")
	builder.WriteString(d.wanpai.String())
	builder.WriteString("\n")
	for _, p := range d.players {
		builder.WriteString(p.String())
		builder.WriteString("\n")
	}
	return builder.String()
}

// 初始化对局
// 初始化玩家、场况信息
func (d *Dealer) Init() {
	for i := 0; i < int(*d.config.Players); i++ {
		d.players = append(d.players, player.New(common.Fon(i+1), *d.config.Tenbo, *d.config.Players))
	}
	d.ken = common.Fon_FON_TON
	d.round = 1
}

// 场名，`金之间·四人南`
func (d *Dealer) Name() string {
	return d.name
}

// 圈名，`东3局`
func (d *Dealer) RoundDesc() string {
	writer := strings.Builder{}
	writer.WriteString(util.Fon2String(d.ken))
	writer.WriteString(strconv.Itoa(d.round))
	writer.WriteString("局")
	return writer.String()
}

// 移动当前玩家指示器
func (d *Dealer) move() {
	d.current = (d.current + 1) % (*d.config.Players)
}

// 开局，配牌，重置计时器
func (d *Dealer) NewRound() {
	// 提取一副牌
	d.pack = NewPack(d.config)
	// 开门
	d.wanpai = deck.NewWanpai(d.pack, *d.config.Players)
	// 定位到庄家
	for i, p := range d.players {
		if p.Seat() == common.Fon_FON_TON {
			d.current = int64(i)
			break
		}
	}
	// 抓3组12张
	for i := 0; i < 3; i++ {
		for j := 0; j < len(d.players); j++ {
			for k := 0; k < 4; k++ {
				d.players[d.current].Draw(d.pack)
			}
			d.move()
		}
	}
	// 抓单张
	for i := 0; i < len(d.players); i++ {
		d.players[d.current].Draw(d.pack)
		d.move()
	}
	// 东家抓牌开打
	d.players[d.current].Draw(d.pack)
	// 记录当前牌山
	d.shortState = d.pack.ShortString() + d.wanpai.ShortString()
	d.md5 = util.MD5Sum(d.shortState)

	d.last = d.pack.Front()
}

func (d *Dealer) ShortState() string {
	return d.shortState
}

func (d *Dealer) MD5() string {
	return d.md5
}
