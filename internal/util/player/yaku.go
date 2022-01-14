package player

import (
	"github.com/wzyjerry/mahjong/internal/ent/schema/common"
	"github.com/wzyjerry/mahjong/internal/util"
)

var (
	Richi = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("立直"),
		Desc:      util.PString("门前清状态听牌即可立直，立直状态下和牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_RICHI),
	}
	Tanyao = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("断幺九"),
		Desc:      util.PString("手牌中不包含幺九牌（19万，19筒，19条，字牌）"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_TANYAO),
	}
	Tsumo = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("门前清自摸和"),
		Desc:      util.PString("门前清状态下自摸和牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_TSUMO),
	}
	Menfon = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("役牌：自风牌"),
		Desc:      util.PString("包含自风刻子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_MENFON),
	}
	Chanfon = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("役牌：场风牌"),
		Desc:      util.PString("包含场风刻子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_CHANFON),
	}
	Sangen = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("役牌：三元牌"),
		Desc:      util.PString("包含白、发、中的刻子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_SANGEN),
	}
	Pinfu = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("平和"),
		Desc:      util.PString("4组顺子+非役牌的雀头+最后是顺子的两面听"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_PINFU),
	}
	Ipeko = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("一杯口"),
		Desc:      util.PString("2组完全相同的顺子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_IPEKO),
	}
	Chankan = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("抢杠"),
		Desc:      util.PString("别家加杠的时候荣和（国士无双可以抢暗杠"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_CHANKAN),
	}
	RinshanKaiho = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("岭上开花"),
		Desc:      util.PString("用摸到的岭上牌和牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_RINSHAN_KAIHO),
	}
	Haiteimoyue = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("海底摸月"),
		Desc:      util.PString("最后一张牌自摸和牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_HAITEI_MOYUE),
	}
	Houteiraoyui = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("河底摸鱼"),
		Desc:      util.PString("最后一张牌荣和"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_HOUTEI_RAOYUI),
	}
	Ippatsu = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("一发"),
		Desc:      util.PString("立直后，无人鸣牌的状态下一巡内和牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_IPPATSU),
	}
	Dora = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(false),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("宝牌"),
		Desc:      util.PString("宝牌指示牌的下一张牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_DORA),
	}
	RiDora = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(false),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("里宝牌"),
		Desc:      util.PString("立直和牌时，里宝牌指示牌的下一张牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_RI_DORA),
	}
	AkaDora = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(false),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("赤宝牌"),
		Desc:      util.PString("红5万，红5筒，红5索"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_AKA_DORA),
	}
	PeDora = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(false),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("北宝牌"),
		Desc:      util.PString("在三人麻将中，北风在拔北操作后可以北当作宝牌（手牌中不算）"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_PE_DORA),
	}
	Tsubamegaeshi = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 燕返"),
		Desc:      util.PString("荣和别家的立直宣言牌（仅第一张的宣言牌）"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_TSUBAMEGAESHI),
	}
	Kanfuri = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 杠振"),
		Desc:      util.PString("在别家杠完后打出一张牌时自家荣和"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_KANFURI),
	}
	Shiaruraotai = &common.Yaku{
		Fan:       util.PInt64(1),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 十二落抬"),
		Desc:      util.PString("四副露自摸或者荣和"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_SHIARURAOTAI),
	}
	// 二番
	WRichi = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("两立直"),
		Desc:      util.PString("轮到自己之前无人鸣牌的状态下第一巡就立直"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_WRICHI),
	}
	SanshokuDoko = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("三色同刻"),
		Desc:      util.PString("万，筒，索都有相同数字的刻子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_SANSHOKU_DOKO),
	}
	Sankantsu = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("三杠子"),
		Desc:      util.PString("一人开杠3次"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_SANKANTSU),
	}
	Toitoiho = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("对对和"),
		Desc:      util.PString("拥有4组刻子或者杠"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_TOITOIHO),
	}
	Sananko = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("三暗刻"),
		Desc:      util.PString("拥有3组没有碰的刻子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_SANANKO),
	}
	Shosangen = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("小三元"),
		Desc:      util.PString("包含白、发、中其中2种的刻子+剩下1种的雀头"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_SHOSANGEN),
	}
	Honroutou = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("混老头"),
		Desc:      util.PString("和牌时只包含老头牌（19万，19筒，19索）和字牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_HONROUTOU),
	}
	Chitoitsu = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("七对子"),
		Desc:      util.PString("7组不同的对子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_CHITOITSU),
	}
	Chanta = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(true),
		Name:      util.PString("混全带幺九"),
		Desc:      util.PString("包含老头牌加上字牌的4组顺子和刻子+幺九牌的雀头"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_CHANTA),
	}
	Ikkitsuukan = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(true),
		Name:      util.PString("一气通贯"),
		Desc:      util.PString("同种数牌组成123，456，789的顺子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_IKKITSUUKAN),
	}
	SanshokuDoujun = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(true),
		Name:      util.PString("三色同顺"),
		Desc:      util.PString("万，筒，索都有相同数字的顺子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_SANSHOKU_DOUJUN),
	}
	Umensai = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 五门齐"),
		Desc:      util.PString("包含万，筒，索，风牌，三元牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_UMENSAI),
	}
	Sanrenko = &common.Yaku{
		Fan:       util.PInt64(2),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 三连刻"),
		Desc:      util.PString("包含三个连续的同种类的刻子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_SANRENKO),
	}
	// 三番
	Ryanpeko = &common.Yaku{
		Fan:       util.PInt64(3),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("二杯口"),
		Desc:      util.PString("包含2组一杯口"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_RYANPEKO),
	}
	Junchan = &common.Yaku{
		Fan:       util.PInt64(3),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(true),
		Name:      util.PString("纯全带幺九"),
		Desc:      util.PString("只包含老头牌的4组顺子和刻子+老头牌的雀头"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_JUNCHAN),
	}
	Honiso = &common.Yaku{
		Fan:       util.PInt64(3),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(true),
		Name:      util.PString("混一色"),
		Desc:      util.PString("只包含1种数牌，并且含有字牌的刻子或者雀头"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_HONISO),
	}
	Isshokusanjun = &common.Yaku{
		Fan:       util.PInt64(3),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(true),
		Name:      util.PString("古役: 一色三同顺"),
		Desc:      util.PString("包含三组完全相同的顺子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_ISSHOKUSANJUN),
	}
	// 六番
	Chiniso = &common.Yaku{
		Fan:       util.PInt64(6),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(true),
		Name:      util.PString("清一色"),
		Desc:      util.PString("只包含1种数牌，不能含有字牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_CHINISO),
	}
	// 满贯
	NagashiMangan = &common.Yaku{
		Fan:       util.PInt64(5),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("流局满贯"),
		Desc:      util.PString("自己的舍张全是幺九牌并且没有被他家吃碰杠的状态下荒牌流局"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_NAGASHI_MANGAN),
	}
	Ipinmoyue = &common.Yaku{
		Fan:       util.PInt64(5),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 一筒摸月"),
		Desc:      util.PString("最后一张牌自摸和牌，且这张牌为1筒（五番）"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_IPINMOYUE),
	}
	Chiyupinraoyui = &common.Yaku{
		Fan:       util.PInt64(5),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 九筒捞鱼"),
		Desc:      util.PString("最后一张牌荣和，且这张牌为9筒（五番）"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_CHIYUPINRAOYUI),
	}
	// 役满
	Tenho = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("天和"),
		Desc:      util.PString("庄家第一巡和牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_TENHO),
	}
	Chiho = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("地和"),
		Desc:      util.PString("轮到自己前无人鸣牌的状态下第一巡自摸和牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_CHIHO),
	}
	Daisangen = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("大三元"),
		Desc:      util.PString("包含白、发、中的三组刻子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_DAISANGEN),
	}
	Suankou = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("四暗刻"),
		Desc:      util.PString("包含没有碰的4组刻子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_SUANKOU),
	}
	Tsuiso = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("字一色"),
		Desc:      util.PString("只包含字牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_TSUISO),
	}
	Ryuhiso = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("绿一色"),
		Desc:      util.PString("只包含索子的23468以及发"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_RYUISO),
	}
	Chinroutou = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("清老头"),
		Desc:      util.PString("手牌中只有老头牌（19万，19筒，19索）"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_CHINROUTOU),
	}
	Kokushimusou = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("国士无双"),
		Desc:      util.PString("全部十三种幺九牌各1张外加其中一种再有1张"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_KOKUSHIMUSOU),
	}
	Shaosushi = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("小四喜"),
		Desc:      util.PString("包含三种风牌的刻子+剩下一种风牌的雀头"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_SHAOSUSHI),
	}
	Sukantsu = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("四缸子"),
		Desc:      util.PString("1人开杠4次"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_SUKANTSU),
	}
	Chuurenpoutou = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("九莲宝灯"),
		Desc:      util.PString("同种数牌1112345678999+其中任意一种再有1张"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_CHUURENPOUTOU),
	}
	Renho = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 人和"),
		Desc:      util.PString("身为子家，第一巡无人鸣牌的情况下，轮到自己摸牌前荣和"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_RENHO),
	}
	Daisharin = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 大车轮"),
		Desc:      util.PString("由筒子“2-8”各两张组成的七对子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_DAISHARIN),
	}
	Daichikurin = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 大竹林"),
		Desc:      util.PString("由索子“2-8”各两张组成的七对子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_DAICHIKURIN),
	}
	Daisuurin = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 大数邻"),
		Desc:      util.PString("由万子“2-8”各两张组成的七对子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_DAISUURIN),
	}
	IshigamiSannen = &common.Yaku{
		Fan:       util.PInt64(16),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 石上三年"),
		Desc:      util.PString("两立直+海底捞月 或 两立直+河堤捞鱼"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_ISHIGAMI_SANNEN),
	}
	// 双倍役满
	SuankouTanki = &common.Yaku{
		Fan:       util.PInt64(32),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("四暗刻单骑"),
		Desc:      util.PString("四暗刻最后单骑听牌和牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_SUANKOU_TANKI),
	}
	KokushimusouJusanmen = &common.Yaku{
		Fan:       util.PInt64(32),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("国士无双十三面"),
		Desc:      util.PString("国士无双最后13面听牌和牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_KOKUSHIMUSOU_JUSANMEN),
	}
	JunseiChuurenpoutou = &common.Yaku{
		Fan:       util.PInt64(32),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("纯正九莲宝灯"),
		Desc:      util.PString("九莲宝灯最后9面听牌和牌"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_JUNSEI_CHUURENPOUTOU),
	}
	Daisushi = &common.Yaku{
		Fan:       util.PInt64(32),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(false),
		Kuisagari: util.PBool(false),
		Name:      util.PString("大四喜"),
		Desc:      util.PString("包含4种风牌的刻子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_DAISUSHI),
	}
	Daichishin = &common.Yaku{
		Fan:       util.PInt64(32),
		Yaku:      util.PBool(true),
		Koyaku:    util.PBool(true),
		Kuisagari: util.PBool(false),
		Name:      util.PString("古役: 大七星"),
		Desc:      util.PString("由七种字牌所组成的七对子"),
		YakuKind:  util.PYakuKind(common.YakuKind_YAKU_KIND_DAICHISHIN),
	}
)
