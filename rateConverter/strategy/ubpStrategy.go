package strategy

import (
	"fmt"
	"math"
	"strconv"

	"chillcy.com/rate-converter/global"
	"chillcy.com/rate-converter/utils"
)

// @author: [Gaayean]
// @description: 英镑换算策略
type UBPstrategy struct {
	Price float64
}

func (u *UBPstrategy) ConverterToCNY() (cnyPrice float64, err error) {
	cnyRate, err := utils.GetRateToCNY(global.RATE_CONFIG.RateApi, "UBP")
	if err != nil {
		return
	}
	result := math.Floor(u.Price*cnyRate*100) / 100
	cnyPrice, err = strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
	return
}
