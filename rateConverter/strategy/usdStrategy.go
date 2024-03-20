package strategy

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Biely/chillcy/global"
	"github.com/Biely/chillcy/utils"
)

// @author: [Gaayean]
// @description: 美元换算策略
type USDstrategy struct {
	Price float64
}

func (u *USDstrategy) ConverterToCNY() (cnyPrice float64, err error) {
	cnyRate, err := utils.GetRateToCNY(global.RATE_CONFIG.RateApi, "USD")
	if err != nil {
		return
	}
	result := math.Floor(u.Price*cnyRate*100) / 100
	cnyPrice, err = strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
	return
}
