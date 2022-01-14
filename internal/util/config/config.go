// 封装游戏配置，定义默认配置和类型检查
package config

import (
	"fmt"
	"github.com/wzyjerry/mahjong/internal/ent/schema/common"
	"github.com/wzyjerry/mahjong/internal/util"
	"strconv"
	"strings"
)

type Config struct {
	*common.Config
}

var Default = &Config{
	Config: &common.Config{
		Players:     util.PInt64(4),
		Rounds:      util.PInt64(2),
		Time:        util.PString("5+20"),
		Transparent: util.PBool(false),
		Tenbo:       util.PInt64(25000),
		Need:        util.PInt64(30000),
		Chi:         util.PInt64(3),
		Kuigae:      util.PBool(true),
		Kuitan:      util.PBool(true),
		Koyaku:      util.PBool(false),
		MinFan:      util.PInt64(1),
	},
}

func New(config *common.Config) *Config {
	return &Config{
		Config: config,
	}
}

func (c *Config) String() string {
	builder := strings.Builder{}
	errs := c.Validate()
	if len(errs) > 0 {
		for _, err := range errs {
			builder.WriteString(err.Error())
			builder.WriteString("\n")
		}
	} else {
		switch *c.Players {
		case 3:
			builder.WriteString("三人")
		case 4:
			builder.WriteString("四人")
		}
		if *c.Transparent {
			builder.WriteString("透明")
		}
		switch *c.Rounds {
		case 1:
			builder.WriteString("东")
		case 2:
			builder.WriteString("南")
		case 4:
			builder.WriteString("一局")
		}
		builder.WriteString("\n")
		builder.WriteString("时间")
		builder.WriteString(*c.Time)
		builder.WriteString("\n")
		builder.WriteString("点数")
		builder.WriteString(strconv.Itoa(int(*c.Tenbo)))
		builder.WriteString("/")
		builder.WriteString(strconv.Itoa(int(*c.Need)))
		builder.WriteString("\n")
		builder.WriteString(strconv.Itoa(int(*c.Chi)))
		builder.WriteString("赤")
		builder.WriteString(strconv.Itoa(int(*c.MinFan)))
		builder.WriteString("番缚\n")
		if *c.Kuigae {
			builder.WriteString("有")
		} else {
			builder.WriteString("无")
		}
		builder.WriteString("食替 ")
		if *c.Kuitan {
			builder.WriteString("有")
		} else {
			builder.WriteString("无")
		}
		builder.WriteString("食断 ")
		if *c.Koyaku {
			builder.WriteString("有")
		} else {
			builder.WriteString("无")
		}
		builder.WriteString("古役")
		builder.WriteString("\n")
	}
	return builder.String()
}

func (c *Config) Validate() []error {
	result := make([]error, 0)
	if c == nil || c.Config == nil {
		result = append(result, fmt.Errorf("配置为空"))
		return result
	}
	if c.Players == nil || !(*c.Players == 3 || *c.Players == 4) {
		result = append(result, fmt.Errorf("玩家数应为3或4"))
	}
	if c.Rounds == nil || !(*c.Rounds == 1 || *c.Rounds == 2 || *c.Rounds == 4) {
		result = append(result, fmt.Errorf("局数应为1、2或4"))
	}
	if c.Time == nil {
		result = append(result, fmt.Errorf("时间应为(3+5/5+10/5+20/60+0/300+0)"))
	} else {
		time := strings.Split(*c.Time, "+")
		if len(time) != 2 {
			result = append(result, fmt.Errorf("时间应为(3+5/5+10/5+20/60+0/300+0)"))
		} else {
			a, b := time[0], time[1]
			if !((a == "3" && b == "5") || (a == "5" && b == "10") || (a == "5" && b == "20") || (a == "60" && b == "0") || (a == "300" && b == "0")) {
				result = append(result, fmt.Errorf("时间应为(3+5/5+10/5+20/60+0/300+0)"))
			}
		}
	}
	if c.Transparent == nil {
		result = append(result, fmt.Errorf("应设置是否透明"))
	}
	if c.Tenbo == nil {
		result = append(result, fmt.Errorf("应设置起始点数"))
	} else {
		if *c.Tenbo%100 != 0 || !(0 <= *c.Tenbo && *c.Tenbo <= 200000) {
			result = append(result, fmt.Errorf("起始点数最小单位100, 最大值200000"))
		}
	}
	if c.Need == nil {
		result = append(result, fmt.Errorf("应设置1位必要点数"))
	} else {
		if *c.Need%100 != 0 || !(0 <= *c.Need && *c.Need <= 200000) {
			result = append(result, fmt.Errorf("1位必要点数最小单位100, 最大值200000"))
		}
		if c.Tenbo != nil && *c.Need <= *c.Tenbo {
			result = append(result, fmt.Errorf("1位必要点数必须大于起始点数"))
		}
	}
	if c.Chi == nil || !(*c.Chi == 0 || *c.Chi == 3 || *c.Chi == 4) {
		result = append(result, fmt.Errorf("赤宝牌数应为0、3或4"))
	}
	if c.Kuigae == nil {
		result = append(result, fmt.Errorf("应设置是否食断"))
	}
	if c.Kuitan == nil {
		result = append(result, fmt.Errorf("应设置是否食断"))
	}
	if c.Koyaku == nil {
		result = append(result, fmt.Errorf("应设置是否允许古役"))
	}
	if c.MinFan == nil || !(*c.MinFan == 1 || *c.MinFan == 2 || *c.MinFan == 4) {
		result = append(result, fmt.Errorf("番缚应为1、2或4"))
	}
	return result
}
