package strategy

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Biely/chillcy/global"
	"github.com/Biely/chillcy/utils"
)

// @author: [Gaayean]
// @description: 欧元换算策略
type EURstrategy struct {
	Price float64
}

func (e *EURstrategy) ConverterToCNY() (cnyPrice float64, err error) {
	cnyRate, err := utils.GetRateToCNY(global.RATE_CONFIG.RateApi, "EUR")
	if err != nil {
		return
	}
	result := math.Floor(e.Price*cnyRate*100) / 100
	cnyPrice, err = strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
	return
}
