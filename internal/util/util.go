package util

import (
	"crypto/md5"
	"fmt"
	"github.com/wzyjerry/mahjong/internal/ent/schema/common"
	"sort"
	"strconv"
	"strings"
)

var cryptoMD5 = md5.New()

func PInt64(i int64) *int64 {
	return &i
}

func PBool(b bool) *bool {
	return &b
}

func PString(s string) *string {
	return &s
}

func PSuit(s common.Suit) *common.Suit {
	return &s
}

func PFon(f common.Fon) *common.Fon {
	return &f
}

func PYakuKind(y common.YakuKind) *common.YakuKind {
	return &y
}

func MD5Sum(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

type mpsz struct {
	number float64
	suit   string
}

func Fon2String(f common.Fon) string {
	switch f {
	case common.Fon_FON_TON:
		return "东"
	case common.Fon_FON_NAN:
		return "南"
	case common.Fon_FON_SHA:
		return "西"
	case common.Fon_FON_PE:
		return "北"
	default:
		return "未知"
	}
}

// 将mpsz表示排序并压缩，如果有14张牌，取最后一张作为进张
func MPSZCompress(str string) string {
	if len(str) == 0 || len(str)%2 != 0 {
		return "错误的mpsz串"
	}
	suffix := ""
	if len(str) == 28 {
		suffix = str[26:28]
		str = str[:26]
	}
	lst := make([]mpsz, 0, len(str)>>1)
	for i := 0; i < len(str); i += 2 {
		f, _ := strconv.ParseFloat(str[i:i+1], 64)
		if f == 0 {
			f = 4.5
		}
		lst = append(lst, mpsz{
			number: f,
			suit:   str[i+1 : i+2],
		})
	}
	sort.SliceStable(lst, func(i, j int) bool {
		if lst[i].suit == lst[j].suit {
			return lst[i].number < lst[j].number
		}
		return lst[i].suit < lst[j].suit
	})
	builder := strings.Builder{}
	current := lst[0].suit
	for _, item := range lst {
		if current != item.suit {
			builder.WriteString(current)
			current = item.suit
		}
		if item.number == 4.5 {
			builder.WriteString("0")
		} else {
			builder.WriteString(strconv.Itoa(int(item.number)))
		}
	}
	builder.WriteString(current)
	builder.WriteString(suffix)
	return builder.String()
}
