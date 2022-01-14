// 将tokenNested.Tile私有化并添加前后指针，作为双向循环链表的节点
// root节点为tile==nil的特殊点
package deck

import (
	"github.com/wzyjerry/mahjong/internal/ent/schema/common"
	"strconv"
)

type Tile struct {
	*common.Tile

	next, prev *Tile
	deck       *Deck
}

func NewTile(raw *common.Tile) *Tile {
	return &Tile{
		Tile: raw,
	}
}

func (t *Tile) IsRoot() bool {
	return t.Tile == nil
}

func (t *Tile) Next() *Tile {
	return t.next
}

func (t *Tile) Prev() *Tile {
	return t.prev
}

func (t *Tile) Equal(t2 *Tile) bool {
	return t.GetSuit() == t2.GetSuit() && t.GetNumber() == t2.GetNumber()
}

func (t *Tile) NextTo(t2 *Tile) bool {
	return t.GetSuit() != common.Suit_SUIT_TSUPAI && t.GetSuit() == t2.GetSuit() && t.GetNumber()-1 == t2.GetNumber()
}

func (t *Tile) Less(t2 *Tile) bool {
	if t.GetSuit() == t2.GetSuit() {
		return t.GetNumber() < t2.GetNumber()
	}
	return t.GetSuit() < t2.GetSuit()
}

func (t *Tile) String() string {
	number := strconv.Itoa(int(t.GetNumber()))
	prefix := ""
	if t.GetTransparent() {
		prefix = "明"
	}
	if t.GetChi() {
		prefix += "赤"
	}
	switch t.GetSuit() {
	case common.Suit_SUIT_MANZU:
		return prefix + number + "万"
	case common.Suit_SUIT_PINZU:
		return prefix + number + "筒"
	case common.Suit_SUIT_SOUZU:
		return prefix + number + "索"
	case common.Suit_SUIT_TSUPAI:
		switch t.GetNumber() {
		case 1:
			return prefix + "东"
		case 2:
			return prefix + "南"
		case 3:
			return prefix + "西"
		case 4:
			return prefix + "北"
		case 5:
			return prefix + "白"
		case 6:
			return prefix + "发"
		case 7:
			return prefix + "中"
		default:
			return "未知字牌"
		}
	default:
		return "未知花色"
	}
}
