// 王牌在牌堆的基础上添加了岭上牌指示器、开杠计数和宝牌、里宝牌相关接口
package deck

import (
	"fmt"
	"strings"
)

type Wanpai struct {
	*Deck
	count     int   // 杠计数
	indicator *Tile // 宝牌指示牌
}

// 开杠，count+1，需要确保count < 4，返回岭上牌
func (w *Wanpai) Kan(pack *Deck) *Tile {
	n := pack.Back()
	pack.Remove(n)
	w.PushFront(n)
	t := w.Back()
	w.Remove(t)
	w.count++
	return t
}

// 拔北，返回岭上牌
func (w *Wanpai) Nuku(pack *Deck) *Tile {
	n := pack.Back()
	pack.Remove(n)
	w.PushFront(n)
	t := w.Back()
	w.Remove(t)
	return t
}

// 开门
func NewWanpai(pack *Deck, players int64) *Wanpai {
	deck := NewDeck(players)
	pos := 4
	if players == 3 {
		pos = 8
	}
	var indicator *Tile
	for i := 0; i < 14; i++ {
		t := pack.Back()
		pack.Remove(t)
		deck.PushFront(t)
		if i == pos {
			indicator = t
		}
	}

	return &Wanpai{
		Deck:      deck,
		count:     0,
		indicator: indicator,
	}
}

// 获取宝牌指示牌，按序返回count+1张指示牌
func (w *Wanpai) GetDoraIndicators() []*Tile {
	result := make([]*Tile, 0, w.count+1)
	for i, t := 0, w.indicator; i <= w.count; i, t = i+1, t.prev.prev {
		result = append(result, t)
	}
	return result
}

// 打印王牌牌堆
func (w *Wanpai) String() string {
	builder := strings.Builder{}
	indicators := make(map[*Tile]struct{})
	innerIndicators := make(map[*Tile]struct{})
	for _, t := range w.GetDoraIndicators() {
		indicators[t] = struct{}{}
		innerIndicators[t.prev] = struct{}{}
	}
	for i, t := 0, w.Front(); !t.IsRoot(); i, t = i+1, t.Next() {
		if i%12 == 0 {
			builder.WriteString("\n")
		} else if i%int(w.players) == 0 {
			builder.WriteString("\t")
		}
		if _, ok := indicators[t]; ok {
			builder.WriteString("宝")
		}
		if _, ok := innerIndicators[t]; ok {
			builder.WriteString("里")
		}
		builder.WriteString(fmt.Sprintf("%s", t))
		builder.WriteString("\t")
	}
	return builder.String()
}
