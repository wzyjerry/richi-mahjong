package main

import (
	"fmt"
	"github.com/wzyjerry/mahjong/internal/ent/schema/common"
	"github.com/wzyjerry/mahjong/internal/util"
	"github.com/wzyjerry/mahjong/internal/util/config"
	"github.com/wzyjerry/mahjong/internal/util/dealer"
	"github.com/wzyjerry/mahjong/internal/util/deck"
	"github.com/wzyjerry/mahjong/internal/util/player"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	config := config.Default
	config.Transparent = util.PBool(true)
	config.Players = util.PInt64(3)
	dealer := dealer.New("金之间", config)
	dealer.Init()
	dealer.NewRound()
	fmt.Println(dealer)

	p := player.New(common.Fon_FON_TON, 35000, 3)
	d := deck.NewDeck(3)

	wanpai := dealer.Wanpai()
	wanpai.GetDoraIndicators()[0].Tile = player.GenTile(common.Suit_SUIT_TSUPAI, 1).Tile
	wanpai.GetDoraIndicators()[0].Prev().Tile = player.GenTile(common.Suit_SUIT_TSUPAI, 2).Tile
	fmt.Println(wanpai)

	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 1))
	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 2))
	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 3))
	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 3))
	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 4))
	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 4))
	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 5))
	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 5))
	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 6))
	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 6))
	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 7))
	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 8))
	//d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 9))

	d.PushBack(player.GenTile(common.Suit_SUIT_MANZU, 2))
	d.PushBack(player.GenTile(common.Suit_SUIT_MANZU, 2))
	d.PushBack(player.GenTile(common.Suit_SUIT_MANZU, 3))
	d.PushBack(player.GenTile(common.Suit_SUIT_MANZU, 4))
	d.PushBack(player.GenTile(common.Suit_SUIT_MANZU, 4))
	d.PushBack(player.GenTile(common.Suit_SUIT_MANZU, 5))
	d.PushBack(player.GenTile(common.Suit_SUIT_MANZU, 5))
	d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 2))
	d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 3))
	d.PushBack(player.GenTile(common.Suit_SUIT_PINZU, 4))
	d.PushBack(player.GenTile(common.Suit_SUIT_SOUZU, 2))
	d.PushBack(player.GenTile(common.Suit_SUIT_SOUZU, 3))
	d.PushBack(player.GenTile(common.Suit_SUIT_SOUZU, 4))

	for i := 0; i < 14; i++ {
		p.Draw(d)
	}
	ron := player.GenTile(common.Suit_SUIT_MANZU, 3)
	p.SetRon(ron)

	// TODO: 番缚

	//p.Set()
	fmt.Printf("%s %s\n", p, ron)
	fan, fu, yakus := p.GetScore(&common.AgariStatus{
		Ken:           util.PFon(common.Fon_FON_TON),
		Kuitan:        util.PBool(true),
		Koyaku:        util.PBool(true),
		MinFan:        util.PInt64(1),
		Tenho:         util.PBool(false),
		Chiho:         util.PBool(false),
		Renho:         util.PBool(false),
		Chankan:       util.PBool(false),
		Rinshankaiho:  util.PBool(false),
		Saigo:         util.PBool(false),
		Tsubamegaeshi: util.PBool(false),
		Kanfuri:       util.PBool(false),
	}, wanpai)
	fmt.Printf("%d番%d符\n", fan, fu)
	for _, yaku := range yakus {
		fmt.Printf("%s %d番\n", yaku.GetName(), yaku.GetFan())
	}
}
