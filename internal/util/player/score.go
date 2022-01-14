package player

import (
	"github.com/wzyjerry/mahjong/internal/ent/schema/common"
	"github.com/wzyjerry/mahjong/internal/util"
	"github.com/wzyjerry/mahjong/internal/util/deck"
	"sort"
)

func tanyao(allTiles []*deck.Tile) bool {
	for _, t := range allTiles {
		if t.GetNumber() == 1 || t.GetNumber() == 9 || t.GetSuit() == common.Suit_SUIT_TSUPAI {
			return false
		}
	}
	return true
}

func countTsupai(allTiles []*deck.Tile) [8]int {
	counter := [8]int{}
	for _, t := range allTiles {
		if t.GetSuit() == common.Suit_SUIT_TSUPAI {
			counter[t.GetNumber()]++
		}
	}
	return counter
}

// 一般型拆分，sorted排序后的手牌（包含铳牌），jantou是否包含雀头，used已使用牌
func (p *player) splitNormal(sorted []*deck.Tile, jantou bool, used *[14]bool, orders []*deck.Ordered, results *[][]*deck.Ordered) {
	if len(orders) == 5 {
		*results = append(*results, orders)
	}
	for i := 0; i < len(sorted); i++ {
		// 找到第一张未使用的牌
		if !used[i] {
			// 作为雀头
			if jantou == false {
				for j := i + 1; j < len(sorted); j++ {
					if !used[j] && sorted[i].Equal(sorted[j]) {
						used[i], used[j] = true, true
						order := deck.NewOrdered()
						order.SetFrom(p.seat)
						order.SetKind(common.OrderedKind_ORDERED_KIND_JANTOU)
						order.PushBack(deck.NewTile(sorted[i].Tile))
						order.PushBack(deck.NewTile(sorted[j].Tile))
						orders = append(orders, order)
						p.splitNormal(sorted, true, used, orders, results)
						orders = orders[:len(orders)-1]
						used[i], used[j] = false, false
					}
				}
			}
			// 作为刻子
			if len(orders) < 5 {
				for j := i + 1; j < len(sorted); j++ {
					if !used[j] && sorted[i].Equal(sorted[j]) {
						for k := j + 1; k < len(sorted); k++ {
							if !used[k] && sorted[j].Equal(sorted[k]) {
								used[i], used[j], used[k] = true, true, true
								order := deck.NewOrdered()
								order.SetFrom(p.seat)
								order.SetKind(common.OrderedKind_ORDERED_KIND_PON)
								order.PushBack(deck.NewTile(sorted[i].Tile))
								order.PushBack(deck.NewTile(sorted[j].Tile))
								order.PushBack(deck.NewTile(sorted[k].Tile))
								orders = append(orders, order)
								p.splitNormal(sorted, jantou, used, orders, results)
								orders = orders[:len(orders)-1]
								used[i], used[j], used[k] = false, false, false
							}
						}
					}
				}
			}
			// 作为顺子
			if len(orders) < 5 && sorted[i].GetSuit() != common.Suit_SUIT_TSUPAI {
				for j := i + 1; j < len(sorted); j++ {
					if !used[j] && sorted[j].NextTo(sorted[i]) {
						for k := j + 1; k < len(sorted); k++ {
							if !used[k] && sorted[k].NextTo(sorted[j]) {
								used[i], used[j], used[k] = true, true, true
								order := deck.NewOrdered()
								order.SetFrom(p.seat)
								order.SetKind(common.OrderedKind_ORDERED_KIND_CHI)
								order.PushBack(deck.NewTile(sorted[i].Tile))
								order.PushBack(deck.NewTile(sorted[j].Tile))
								order.PushBack(deck.NewTile(sorted[k].Tile))
								orders = append(orders, order)
								p.splitNormal(sorted, jantou, used, orders, results)
								orders = orders[:len(orders)-1]
								used[i], used[j], used[k] = false, false, false
							}
						}
					}
				}
			}
		}
	}
}

// 和牌后算分，返回番、符、役种列表
func (p *player) GetScore(status *common.AgariStatus, wanpai *deck.Wanpai) (int64, int64, []*common.Yaku) {
	// allTiles 全部牌
	allTiles := make([]*deck.Tile, 0, 18)
	for _, furo := range p.furoList {
		allTiles = append(allTiles, furo.Tiles()...)
	}
	allTiles = append(allTiles, p.hand.Tiles()...)
	if p.ron != nil {
		allTiles = append(allTiles, p.ron)
	}
	// 最后一张
	lastTile := allTiles[len(allTiles)-1]
	// 排序手牌（包含铳牌）
	sorted := p.hand.Tiles()
	if p.ron != nil {
		sorted = append(sorted, p.ron)
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Less(sorted[j])
	})
	// 手牌拆分
	var results [][]*deck.Ordered
	p.splitNormal(sorted, false, &[14]bool{}, p.furoList, &results)
	var (
		fan, fu int64
		yakus   []*common.Yaku
	)

	if p.isMenzentin() {
		ok, tile := testKokushi(sorted)
		if ok {
			fu = 25
			if tile.Equal(lastTile) {
				fan = KokushimusouJusanmen.GetFan()
				yakus = []*common.Yaku{KokushimusouJusanmen}
			} else {
				fan = Kokushimusou.GetFan()
				yakus = []*common.Yaku{Kokushimusou}
			}
		}
		if testChitoitsu(sorted) {
			fu = 25
			fan, yakus = p.scoreChitoitsu(sorted, status, wanpai)
		}
	}
	// 一般型拆分
	for _, kind := range results {
		subFan, subFu, subYakus := p.scoreNormal(allTiles, kind, lastTile, status, wanpai)
		if subFan == fan {
			if subFu > fu {
				fan, fu, yakus = subFan, subFu, subYakus
			}
		} else if subFan > fan {
			fan, fu, yakus = subFan, subFu, subYakus
		}
	}
	if status.GetTenho() || status.GetChiho() {
		if len(yakus) != 0 && *yakus[0].Fan >= 16 {
			if status.GetTenho() {
				yakus = append([]*common.Yaku{Tenho}, yakus...)
			} else {
				yakus = append([]*common.Yaku{Chiho}, yakus...)
			}
		} else {
			fan = 16
			if status.GetTenho() {
				yakus = []*common.Yaku{Tenho}
			} else {
				yakus = []*common.Yaku{Chiho}
			}
		}
	}
	if status.GetKoyaku() {
		if status.GetRenho() {
			if len(yakus) != 0 && *yakus[0].Fan >= 16 {
				yakus = append([]*common.Yaku{Renho}, yakus...)
			} else {
				fan = *Renho.Fan
				yakus = []*common.Yaku{Renho}
			}
		}
		// 石上三年
		indicator := 0
		for _, yaku := range yakus {
			switch yaku.GetYakuKind() {
			case common.YakuKind_YAKU_KIND_WRICHI:
				indicator |= 1
			case common.YakuKind_YAKU_KIND_HAITEI_MOYUE:
				fallthrough
			case common.YakuKind_YAKU_KIND_HOUTEI_RAOYUI:
				indicator |= 2
			}
		}
		if indicator == 3 {
			fan = IshigamiSannen.GetFan()
			yakus = []*common.Yaku{IshigamiSannen}
		}
	}
	// 非役满
	if status.GetKoyaku() && len(yakus) > 0 && yakus[0].GetFan() < 16 {
		for i, yaku := range yakus {
			switch yaku.GetYakuKind() {
			case common.YakuKind_YAKU_KIND_HAITEI_MOYUE:
				if lastTile.GetSuit() == common.Suit_SUIT_PINZU && lastTile.GetNumber() == 1 {
					yakus = append(yakus[:i], append([]*common.Yaku{Ipinmoyue}, yakus[i+1:]...)...)
					fan += Ipinmoyue.GetFan() - Haiteimoyue.GetFan()
					break
				}
			case common.YakuKind_YAKU_KIND_HOUTEI_RAOYUI:
				if lastTile.GetSuit() == common.Suit_SUIT_PINZU && lastTile.GetNumber() == 9 {
					yakus = append(yakus[:i], append([]*common.Yaku{Chiyupinraoyui}, yakus[i+1:]...)...)
					fan += Chiyupinraoyui.GetFan() - Houteiraoyui.GetFan()
					break
				}
			}
		}
	}
	return fan, fu, yakus
}

func (p *player) scoreNormal(allTiles []*deck.Tile, kind []*deck.Ordered, lastTile *deck.Tile, status *common.AgariStatus, wanpai *deck.Wanpai) (int64, int64, []*common.Yaku) {
	menzentin := p.isMenzentin()
	var yakus []*common.Yaku
	// 四暗刻相关
	ankou := 0
	kan := 0
	pon := 0
	tanki := true
	for _, k := range kind {
		switch k.Kind() {
		case common.OrderedKind_ORDERED_KIND_PON:
			pon++
			if k.From() == p.seat {
				ankou++
				if k.Front() == lastTile {
					tanki = false
				}
			}
		case common.OrderedKind_ORDERED_KIND_KAN:
			kan++
			if k.From() == p.seat {
				ankou++
				if k.Front() == lastTile {
					tanki = false
				}
			}
		}
	}
	if ankou == 4 {
		if tanki {
			yakus = append(yakus, SuankouTanki)
		} else {
			yakus = append(yakus, Suankou)
		}
	}
	if kan == 4 {
		yakus = append(yakus, Sukantsu)
	}
	suits := suits(allTiles)
	if len(suits) == 1 {
		if suits[0] == common.Suit_SUIT_TSUPAI {
			yakus = append(yakus, Tsuiso)
		} else {
			// 处理九莲宝灯相关
			if menzentin {
				counter := [10]int{}
				for _, t := range allTiles {
					counter[t.GetNumber()]++
				}
				accept := true
				for i := 1; i <= 9; i++ {
					if counter[i] == 0 {
						accept = false
						break
					}
				}
				if accept && counter[1] >= 3 && counter[9] >= 3 {
					if counter[lastTile.GetNumber()]&1 == 0 {
						yakus = append(yakus, JunseiChuurenpoutou)
					} else {
						yakus = append(yakus, Chuurenpoutou)
					}
				}
			}
		}
	}
	counter := countTsupai(allTiles)
	if counter[1] >= 3 && counter[2] >= 3 && counter[3] >= 3 && counter[4] >= 3 {
		yakus = append(yakus, Daisushi)
	} else if counter[1] >= 2 && counter[2] >= 2 && counter[3] >= 2 && counter[4] >= 2 {
		yakus = append(yakus, Shaosushi)
	}
	if counter[5] >= 3 && counter[6] >= 3 && counter[7] >= 3 {
		yakus = append(yakus, Daisangen)
	}
	// 绿一色
	accept := true
RyuhisoOuter:
	for _, a := range kind {
		switch a.Kind() {
		case common.OrderedKind_ORDERED_KIND_CHI:
			if !(a.Front().GetSuit() == common.Suit_SUIT_SOUZU && a.Front().GetNumber() == 2) {
				accept = false
				break RyuhisoOuter
			}
		default:
			switch a.Front().GetSuit() {
			case common.Suit_SUIT_SOUZU:
				if a.Front().GetNumber() == 1 || a.Front().GetNumber() == 5 || a.Front().GetNumber() == 7 || a.Front().GetNumber() == 9 {
					accept = false
					break RyuhisoOuter
				}
			case common.Suit_SUIT_TSUPAI:
				if a.Front().GetNumber() != 6 {
					accept = false
					break RyuhisoOuter
				}
			default:
				accept = false
				break RyuhisoOuter
			}
		}
	}
	if accept {
		yakus = append(yakus, Ryuhiso)
	}
	allRoutou := allRoutou(allTiles)
	if allRoutou && suits[0] != common.Suit_SUIT_TSUPAI {
		yakus = append(yakus, Chinroutou)
	}
	// 非役满
	if len(yakus) == 0 {
		if len(suits) == 1 {
			yakus = append(yakus, Chiniso)
		}
		chiMap := make(map[int64]int)
		ponMap := make(map[int64]int)
		for _, k := range kind {
			switch k.Kind() {
			case common.OrderedKind_ORDERED_KIND_CHI:
				chiMap[p.mappingKey(k.Front().GetSuit(), k.Front().GetNumber())]++
			case common.OrderedKind_ORDERED_KIND_PON:
				fallthrough
			case common.OrderedKind_ORDERED_KIND_KAN:
				ponMap[p.mappingKey(k.Front().GetSuit(), k.Front().GetNumber())]++
			}
		}
		if menzentin {
			count := 0
			for _, num := range chiMap {
				if num&1 == 0 {
					count += num >> 1
				}
			}
			if count == 2 {
				yakus = append(yakus, Ryanpeko)
			} else if count == 1 {
				yakus = append(yakus, Ipeko)
			}
			// 平和
			if p.pinfu(kind, lastTile, status) {
				yakus = append(yakus, Pinfu)
			}
		}
		// 全带
		accept := true
	JunchanOuter:
		for _, k := range kind {
			switch k.Kind() {
			case common.OrderedKind_ORDERED_KIND_CHI:
				if k.Front().GetNumber() != 1 && k.Front().GetNumber() != 7 {
					accept = false
					break JunchanOuter
				}
			default:
				if k.Front().GetSuit() != common.Suit_SUIT_TSUPAI && k.Front().GetNumber() != 1 && k.Front().GetNumber() != 9 {
					accept = false
					break JunchanOuter
				}
			}
		}
		if accept {
			if suits[0] != common.Suit_SUIT_TSUPAI {
				// 纯全
				yakus = append(yakus, Junchan)
			} else {
				// 混全
				yakus = append(yakus, Chanta)
			}
		}
		if len(suits) == 2 && suits[0] == common.Suit_SUIT_TSUPAI {
			yakus = append(yakus, Honiso)
		}
		if status.GetKoyaku() {
			for _, num := range chiMap {
				if num >= 3 {
					yakus = append(yakus, Isshokusanjun)
					break
				}
			}
		}
		for _, k := range kind {
			if (k.Kind() == common.OrderedKind_ORDERED_KIND_PON || k.Kind() == common.OrderedKind_ORDERED_KIND_KAN) && k.Front().GetSuit() == common.Suit_SUIT_MANZU {
				if _, ok := ponMap[p.mappingKey(common.Suit_SUIT_PINZU, k.Front().GetNumber())]; ok {
					if _, ok := ponMap[p.mappingKey(common.Suit_SUIT_SOUZU, k.Front().GetNumber())]; ok {
						yakus = append(yakus, SanshokuDoko)
						break
					}
				}
			}
		}
		if kan == 3 {
			yakus = append(yakus, Sankantsu)
		}
		if kan+pon == 4 {
			yakus = append(yakus, Toitoiho)
		}
		if ankou == 3 {
			yakus = append(yakus, Sananko)
		}
		if counter[5] >= 2 && counter[6] >= 2 && counter[7] >= 2 {
			yakus = append(yakus, Shosangen)
		}
		if allRoutou {
			yakus = append(yakus, Honroutou)
		}
		// 一气通贯
		indicator := make(map[common.Suit]int)
		for _, k := range kind {
			if k.Kind() == common.OrderedKind_ORDERED_KIND_CHI {
				if k.Front().GetNumber() == 1 {
					indicator[k.Front().GetSuit()] |= 1
				} else if k.Front().GetNumber() == 4 {
					indicator[k.Front().GetSuit()] |= 2
				} else if k.Front().GetNumber() == 7 {
					indicator[k.Front().GetSuit()] |= 4
				}
			}
		}
		for _, num := range indicator {
			if num == 7 {
				yakus = append(yakus, Ikkitsuukan)
				break
			}
		}
		for _, k := range kind {
			if k.Kind() == common.OrderedKind_ORDERED_KIND_CHI && k.Front().GetSuit() == common.Suit_SUIT_MANZU {
				if _, ok := chiMap[p.mappingKey(common.Suit_SUIT_PINZU, k.Front().GetNumber())]; ok {
					if _, ok := chiMap[p.mappingKey(common.Suit_SUIT_SOUZU, k.Front().GetNumber())]; ok {
						yakus = append(yakus, SanshokuDoujun)
						break
					}
				}
			}
		}
		if status.GetKoyaku() {
			indicator := 0
			for _, k := range kind {
				indicator |= 1 << k.Front().GetSuit()
			}
			if indicator == 31 {
				yakus = append(yakus, Umensai)
			}
			for _, k := range kind {
				if k.Kind() == common.OrderedKind_ORDERED_KIND_PON || k.Kind() == common.OrderedKind_ORDERED_KIND_KAN {
					if k.Front().GetSuit() != common.Suit_SUIT_TSUPAI && k.Front().GetNumber() <= 7 {
						if _, ok := ponMap[p.mappingKey(k.Front().GetSuit(), k.Front().GetNumber()+1)]; ok {
							if _, ok := ponMap[p.mappingKey(k.Front().GetSuit(), k.Front().GetNumber()+2)]; ok {
								yakus = append(yakus, Sanrenko)
								break
							}
						}
					}
				}
			}
			count := 0
			for _, k := range kind {
				if k.From() != p.seat {
					count++
				}
			}
			if count == 4 {
				yakus = append(yakus, Shiaruraotai)
			}
		}
		if counter[p.seat] >= 3 {
			yakus = append(yakus, Menfon)
		}
		if counter[status.GetKen()] >= 3 {
			yakus = append(yakus, Chanfon)
		}
		if counter[5] >= 3 || counter[6] >= 3 || counter[7] >= 3 {
			yakus = append(yakus, Sangen)
		}
		yakus = append(yakus, p.commonYaku(allTiles, wanpai, status)...)
	}

	var fan int64
	for _, yaku := range yakus {
		if !menzentin && yaku.GetKuisagari() {
			fan += yaku.GetFan() - 1
		} else {
			fan += yaku.GetFan()
		}
	}
	return fan, p.calcFu(kind, lastTile, status), yakus
}

func (p *player) pinfu(kind []*deck.Ordered, lastTile *deck.Tile, status *common.AgariStatus) bool {
	for _, k := range kind {
		switch k.Kind() {
		case common.OrderedKind_ORDERED_KIND_CHI:
			if k.Front().Next().GetName() == lastTile.GetName() || (k.Front().GetName() == lastTile.GetName() && lastTile.GetNumber() == 7) || (k.Back().GetName() == lastTile.GetName() && lastTile.GetNumber() == 3) {
				return false
			}
		case common.OrderedKind_ORDERED_KIND_PON:
			return false
		case common.OrderedKind_ORDERED_KIND_KAN:
			return false
		case common.OrderedKind_ORDERED_KIND_JANTOU:
			if k.Front().GetSuit() == common.Suit_SUIT_TSUPAI && (k.Front().GetNumber() == int64(p.seat) || k.Front().GetNumber() == int64(status.GetKen()) || k.Front().GetNumber() >= 5) || k.Front().GetName() == lastTile.GetName() || k.Front().Next().GetName() == lastTile.GetName() {
				return false
			}
		}
	}
	return true
}

func (p *player) calcFu(kind []*deck.Ordered, lastTile *deck.Tile, status *common.AgariStatus) int64 {
	menzentin := p.isMenzentin()
	var fu int64 = 20
	pinfu := true
	for _, k := range kind {
		switch k.Kind() {
		case common.OrderedKind_ORDERED_KIND_CHI:
			if k.Front().Next().GetName() == lastTile.GetName() || (k.Front().GetName() == lastTile.GetName() && lastTile.GetNumber() == 7) || (k.Back().GetName() == lastTile.GetName() && lastTile.GetNumber() == 3) {
				pinfu = false
				fu += 2
			}
		case common.OrderedKind_ORDERED_KIND_PON:
			pinfu = false
			d := 0
			if k.Front().IsYaochu() {
				d = 1
			}
			if k.From() != p.seat {
				fu += 2 << d
			} else {
				fu += 4 << d
			}
		case common.OrderedKind_ORDERED_KIND_KAN:
			pinfu = false
			d := 0
			if k.Front().IsYaochu() {
				d = 1
			}
			if k.From() != p.seat {
				fu += 8 << d
			} else {
				fu += 16 << d
			}
		case common.OrderedKind_ORDERED_KIND_JANTOU:
			if k.Front().GetSuit() == common.Suit_SUIT_TSUPAI && (k.Front().GetNumber() == int64(p.seat) || k.Front().GetNumber() == int64(status.GetKen()) || k.Front().GetNumber() >= 5) || k.Front().GetName() == lastTile.GetName() || k.Front().Next().GetName() == lastTile.GetName() {
				pinfu = false
				fu += 2
			}
		}
	}
	if pinfu {
		if !menzentin {
			return 30
		} else {
			if p.ron == nil {
				return 20
			}
			return 30
		}
	}
	if menzentin {
		if p.ron == nil {
			fu += 2
		}
		fu += 10
	} else {
		if p.ron == nil {
			fu += 2
		}
	}
	if fu%10 != 0 {
		fu = (fu/10 + 1) * 10
	}
	return fu
}

// 返回手牌排序后的花色种类，降序排序
func suits(allTiles []*deck.Tile) []common.Suit {
	suitSet := make(map[common.Suit]struct{})
	for _, t := range allTiles {
		suitSet[t.GetSuit()] = struct{}{}
	}
	suits := make([]common.Suit, 0, len(suitSet))
	for suit := range suitSet {
		suits = append(suits, suit)
	}
	sort.Slice(suits, func(i, j int) bool {
		return suits[i] > suits[j]
	})
	return suits
}

// 除了字牌是否都为老头牌
func allRoutou(allTiles []*deck.Tile) bool {
	for _, t := range allTiles {
		if t.GetSuit() != common.Suit_SUIT_TSUPAI && !t.IsYaochu() {
			return false
		}
	}
	return true
}

func (p *player) scoreChitoitsu(sorted []*deck.Tile, status *common.AgariStatus, wanpai *deck.Wanpai) (int64, []*common.Yaku) {
	// 仅有断幺九、混老头、清一色、混一色复合，字一色，古役大七星（不计字一色）
	yakus := []*common.Yaku{Chitoitsu}
	suits := suits(sorted)
	if len(suits) == 1 {
		if suits[0] == common.Suit_SUIT_TSUPAI {
			if status.GetKoyaku() {
				return Daichishin.GetFan(), []*common.Yaku{Daichishin}
			} else {
				return Tsuiso.GetFan(), []*common.Yaku{Tsuiso}
			}
		} else {
			if status.GetKoyaku() && sorted[0].GetNumber() == 2 && sorted[len(sorted)-1].GetNumber() == 8 {
				switch suits[0] {
				case common.Suit_SUIT_MANZU:
					return Daisuurin.GetFan(), []*common.Yaku{Daisuurin}
				case common.Suit_SUIT_PINZU:
					return Daichikurin.GetFan(), []*common.Yaku{Daichikurin}
				case common.Suit_SUIT_SOUZU:
					return Daisharin.GetFan(), []*common.Yaku{Daisharin}
				}
			}
			yakus = append(yakus, Chiniso)
		}
	}
	if len(suits) == 2 && suits[0] == common.Suit_SUIT_TSUPAI {
		yakus = append(yakus, Honiso)
	}
	if allRoutou(sorted) {
		yakus = append(yakus, Honroutou)
	}
	yakus = append(yakus, p.commonYaku(sorted, wanpai, status)...)
	var fan int64
	for _, yaku := range yakus {
		fan += yaku.GetFan()
	}
	return fan, yakus
}

func (p *player) getDora(indicators []*deck.Tile) []*deck.Tile {
	dora := make([]*deck.Tile, 0, len(indicators))
	for _, indicator := range indicators {
		if indicator.GetSuit() == common.Suit_SUIT_MANZU && p.players == 3 {
			if indicator.GetNumber() == 1 {
				dora = append(dora, GenTile(common.Suit_SUIT_MANZU, 9))
			} else {
				dora = append(dora, GenTile(common.Suit_SUIT_MANZU, 1))
			}
		} else if indicator.GetSuit() == common.Suit_SUIT_TSUPAI {
			if indicator.GetNumber() == 4 {
				dora = append(dora, GenTile(common.Suit_SUIT_TSUPAI, 1))
			} else if indicator.GetNumber() == 7 {
				dora = append(dora, GenTile(common.Suit_SUIT_TSUPAI, 5))
			} else {
				dora = append(dora, GenTile(common.Suit_SUIT_TSUPAI, indicator.GetNumber()+1))
			}
		} else {
			if indicator.GetNumber() == 9 {
				dora = append(dora, GenTile(indicator.GetSuit(), 1))
			} else {
				dora = append(dora, GenTile(indicator.GetSuit(), indicator.GetNumber()+1))
			}
		}
	}
	return dora
}

func (p *player) commonYaku(allTiles []*deck.Tile, wanpai *deck.Wanpai, status *common.AgariStatus) []*common.Yaku {
	yakus := make([]*common.Yaku, 0)
	menzentin := p.isMenzentin()
	if status.GetChankan() {
		yakus = append(yakus, Chankan)
	}
	if status.GetRinshankaiho() {
		yakus = append(yakus, RinshanKaiho)
	}
	if status.GetSaigo() {
		if p.ron == nil {
			yakus = append(yakus, Haiteimoyue)
		} else {
			yakus = append(yakus, Houteiraoyui)
		}
	}
	if p.ippatsu {
		yakus = append(yakus, Ippatsu)
	}
	if p.richi {
		yakus = append(yakus, Richi)
	}
	if p.wrichi {
		yakus = append(yakus, WRichi)
	}
	if tanyao(allTiles) && (menzentin || *status.Kuitan) {
		yakus = append(yakus, Tanyao)
	}
	if menzentin && p.ron == nil {
		yakus = append(yakus, Tsumo)
	}
	if status.GetKoyaku() {
		if status.GetTsubamegaeshi() {
			yakus = append(yakus, Tsubamegaeshi)
		}
		if status.GetKanfuri() {
			yakus = append(yakus, Kanfuri)
		}
	}
	indicators := wanpai.GetDoraIndicators()
	doras := p.getDora(indicators)
	var raw int64
	for _, t := range allTiles {
		for _, d := range doras {
			if t.Equal(d) {
				raw++
			}
		}
	}
	if raw > 0 {
		dora := &common.Yaku{
			Fan:       util.PInt64(raw),
			Yaku:      util.PBool(false),
			Koyaku:    util.PBool(false),
			Kuisagari: util.PBool(false),
			Name:      util.PString("宝牌"),
			Desc:      util.PString("宝牌指示牌的下一张牌"),
			YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_DORA),
		}
		yakus = append(yakus, dora)
	}
	if p.richi || p.wrichi {
		innerIndicators := make([]*deck.Tile, 0, len(indicators))
		for _, i := range indicators {
			innerIndicators = append(innerIndicators, i.Prev())
		}
		riDoras := p.getDora(innerIndicators)
		var ri int64
		for _, t := range allTiles {
			for _, d := range riDoras {
				if t.Equal(d) {
					ri++
				}
			}
		}
		if ri > 0 {
			riDora := &common.Yaku{
				Fan:       util.PInt64(ri),
				Yaku:      util.PBool(false),
				Koyaku:    util.PBool(false),
				Kuisagari: util.PBool(false),
				Name:      util.PString("里宝牌"),
				Desc:      util.PString("立直和牌时，里宝牌指示牌的下一张牌"),
				YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_RI_DORA),
			}
			yakus = append(yakus, riDora)
		}
	}
	var aka int64
	for _, t := range allTiles {
		if t.GetChi() {
			aka++
		}
	}
	if aka > 0 {
		akaDora := &common.Yaku{
			Fan:       util.PInt64(aka),
			Yaku:      util.PBool(false),
			Koyaku:    util.PBool(false),
			Kuisagari: util.PBool(false),
			Name:      util.PString("赤宝牌"),
			Desc:      util.PString("红5万，红5筒，红5索"),
			YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_AKA_DORA),
		}
		yakus = append(yakus, akaDora)
	}
	if p.players == 3 && p.nuku > 0 {
		peDora := &common.Yaku{
			Fan:       util.PInt64(p.nuku),
			Yaku:      util.PBool(false),
			Koyaku:    util.PBool(false),
			Kuisagari: util.PBool(false),
			Name:      util.PString("北宝牌"),
			Desc:      util.PString("在三人麻将中，北风在拔北操作后可以北当作宝牌（手牌中不算）"),
			YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_PE_DORA),
		}
		yakus = append(yakus, peDora)
	}
	return yakus
}

func (p *player) isMenzentin() bool {
	menzentin := true
	for _, furo := range p.furoList {
		if furo.From() != p.seat {
			menzentin = false
		}
	}
	return menzentin
}
