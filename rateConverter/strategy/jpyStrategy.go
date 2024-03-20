package strategy

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Biely/chillcy/global"
	"github.com/Biely/chillcy/utils"
)

// @author: [Gaayean]
// @description: 日元换算策略
type JPYstrategy struct {
	Price float64
}

func (j *JPYstrategy) ConverterToCNY() (cnyPrice float64, err error) {
	cnyRate, err := utils.GetRateToCNY(global.RATE_CONFIG.RateApi, "JPY")
	if err != nil {
		return
	}
	result := math.Floor(j.Price*cnyRate*100) / 100
	cnyPrice, err = strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
	return
}
