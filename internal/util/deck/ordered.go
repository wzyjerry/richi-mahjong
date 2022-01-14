// 副露在牌堆基础上添加了副露牌和来源
package deck

import "github.com/wzyjerry/mahjong/internal/ent/schema/common"

type Ordered struct {
	*Deck
	kind common.OrderedKind
	from common.Fon
	tile *Tile
}

func NewOrdered() *Ordered {
	return &Ordered{
		Deck: NewDeck(0),
	}
}

func (f *Ordered) PushBack(t *Tile) *Tile {
	if f.Len() == 0 {
		return f.Deck.PushBack(t)
	}
	if f.Deck.Front().Less(t) {
		return f.Deck.PushBack(t)
	}
	return f.Deck.PushFront(t)
}

func (f *Ordered) From() common.Fon {
	return f.from
}

func (f *Ordered) Tile() *Tile {
	return f.tile
}

func (f *Ordered) Kind() common.OrderedKind {
	return f.kind
}

func (f *Ordered) SetFrom(seat common.Fon) {
	f.from = seat
}

func (f *Ordered) SetTile(t *Tile) {
	f.tile = t
}

func (f *Ordered) SetKind(k common.OrderedKind) {
	f.kind = k
}
