// Deck 使用双向循环链表模拟任意牌堆
// 包括但不限于：手牌、牌山、吃碰牌区、牌河等
// 遍历牌堆
// for e := Deck.Front(); !e.IsRoot(); e = e.Next() {
//     // do something with e
// }
package deck

import (
	"fmt"
	"github.com/wzyjerry/mahjong/internal/ent/schema/common"
	"math/rand"
	"strconv"
	"strings"
)

type Deck struct {
	root    *Tile
	len     int
	players int64
}

// mpsz表示法
func (d *Deck) ShortString() string {
	builder := strings.Builder{}
	for i, t := 0, d.Front(); !t.IsRoot(); i, t = i+1, t.Next() {
		if t.GetChi() {
			builder.WriteString("0")
		} else {
			builder.WriteString(strconv.Itoa(int(t.GetNumber())))
		}
		switch t.GetSuit() {
		case common.Suit_SUIT_MANZU:
			builder.WriteString("m")
		case common.Suit_SUIT_PINZU:
			builder.WriteString("p")
		case common.Suit_SUIT_SOUZU:
			builder.WriteString("s")
		case common.Suit_SUIT_TSUPAI:
			builder.WriteString("z")
		}
	}
	return builder.String()
}

func (d *Deck) String() string {
	builder := strings.Builder{}
	for i, t := 0, d.Front(); !t.IsRoot(); i, t = i+1, t.Next() {
		if i%12 == 0 {
			builder.WriteString("\n")
		} else if i%int(d.players) == 0 {
			builder.WriteString("\t")
		}
		builder.WriteString(fmt.Sprintf("%s", t))
		builder.WriteString("\t")
	}
	return builder.String()
}

func NewDeck(players int64) *Deck {
	return (&Deck{
		players: players,
	}).Init()
}

func (d *Deck) Init() *Deck {
	d.root = NewTile(nil)
	d.root.next = d.root
	d.root.prev = d.root
	d.len = 0
	return d
}

func (d *Deck) Len() int {
	return d.len
}

func (d *Deck) Front() *Tile {
	return d.root.next
}

func (d *Deck) Back() *Tile {
	return d.root.prev
}

// 从牌堆删除t，长度-1，返回t
func (d *Deck) remove(t *Tile) *Tile {
	t.prev.next = t.next
	t.next.prev = t.prev
	t.next = t
	t.prev = t
	t.deck = nil
	d.len--
	return t
}

// 将t插入到at后，长度+1，返回t
func (d *Deck) insert(t, at *Tile) *Tile {
	t.prev = at
	t.next = at.next
	t.prev.next = t
	t.next.prev = t
	t.deck = d
	d.len++
	return t
}

func (d *Deck) Remove(t *Tile) *Tile {
	if t.deck == d {
		d.remove(t)
	}
	return t
}

func (d *Deck) PushFront(t *Tile) *Tile {
	return d.insert(t, d.root)
}

func (d *Deck) PushBack(t *Tile) *Tile {
	return d.insert(t, d.root.prev)
}

func (d *Deck) Shuffle() {
	tiles := make([]*Tile, 0, d.len)
	for d.len > 0 {
		tiles = append(tiles, d.Remove(d.Front()))
	}
	d.Init()
	rand.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})
	for _, t := range tiles {
		d.PushBack(t)
	}
}

func (d *Deck) Tiles() []*Tile {
	tiles := make([]*Tile, 0, d.Len())
	for t := d.Front(); !t.IsRoot(); t = t.Next() {
		tiles = append(tiles, t)
	}
	return tiles
}
