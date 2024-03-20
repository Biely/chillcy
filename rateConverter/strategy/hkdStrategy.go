package strategy

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Biely/chillcy/global"
	"github.com/Biely/chillcy/utils"
)

// @author: [Gaayean]
// @description: 港币换算策略
type HKDstrategy struct {
	Price float64
}

func (h *HKDstrategy) ConverterToCNY() (cnyPrice float64, err error) {
	cnyRate, err := utils.GetRateToCNY(global.RATE_CONFIG.RateApi, "HKD")
	if err != nil {
		return
	}
	result := math.Floor(h.Price*cnyRate*100) / 100
	cnyPrice, err = strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
	return
}
