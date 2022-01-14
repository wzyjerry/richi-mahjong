// Manager为每场对局生成一副新麻将牌
package dealer

import (
	"github.com/wzyjerry/mahjong/internal/ent/schema/common"
	"github.com/wzyjerry/mahjong/internal/util"
	"github.com/wzyjerry/mahjong/internal/util/config"
	"github.com/wzyjerry/mahjong/internal/util/deck"
	"math/rand"
	"strconv"
)

func genTiles(suit common.Suit, number int64, transparent bool, chi int64) []*deck.Tile {
	tiles := make([]*deck.Tile, 0, 4)
	for i := 0; i < 4; i++ {
		isChi := false
		if number == 5 && suit != common.Suit_SUIT_TSUPAI {
			if (chi == 3 && i == 0) || ((chi == 4) && ((i == 0) || (suit == common.Suit_SUIT_PINZU && i == 1))) {
				isChi = true
			}
		}
		tile := &deck.Tile{
			Tile: &common.Tile{
				Suit:        util.PSuit(suit),
				Number:      util.PInt64(number),
				Chi:         util.PBool(isChi),
				Transparent: util.PBool(false),
			},
		}
		tile.Tile.Name = util.PString(tile.String() + "_" + strconv.Itoa(i))
		tiles = append(tiles, tile)
	}
	if transparent {
		none := rand.Intn(4)
		for i := range tiles {
			if i == none {
				tiles[i].Tile.Transparent = util.PBool(false)
			} else {
				tiles[i].Tile.Transparent = util.PBool(true)
			}
		}
	}
	return tiles
}

func NewPack(config *config.Config) *deck.Deck {
	tiles := make([]*deck.Tile, 0, 136)
	// 万
	if *config.Players == 3 {
		tiles = append(tiles, genTiles(common.Suit_SUIT_MANZU, 1, *config.Transparent, *config.Chi)...)
		tiles = append(tiles, genTiles(common.Suit_SUIT_MANZU, 9, *config.Transparent, *config.Chi)...)
	} else {
		for i := int64(1); i <= 9; i++ {
			tiles = append(tiles, genTiles(common.Suit_SUIT_MANZU, i, *config.Transparent, *config.Chi)...)
		}
	}
	// 饼
	for i := int64(1); i <= 9; i++ {
		tiles = append(tiles, genTiles(common.Suit_SUIT_PINZU, i, *config.Transparent, *config.Chi)...)
	}
	// 索
	for i := int64(1); i <= 9; i++ {
		tiles = append(tiles, genTiles(common.Suit_SUIT_SOUZU, i, *config.Transparent, *config.Chi)...)
	}
	// 字
	for i := int64(1); i <= 7; i++ {
		tiles = append(tiles, genTiles(common.Suit_SUIT_TSUPAI, i, *config.Transparent, *config.Chi)...)
	}
	// 初始化牌堆并洗牌
	pack := deck.NewDeck(*config.Players).Init()
	for _, t := range tiles {
		pack.PushBack(t)
	}
	pack.Shuffle()
	return pack
}
