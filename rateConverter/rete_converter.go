package rateconverter

import (
	"fmt"

	"chillcy.com/rate-converter/global"
	"chillcy.com/rate-converter/rateConverter/strategy"
	"chillcy.com/rate-converter/utils"
)

type RateConverter struct {
}

//@author: [Gaayean]
//@function: GetCNY
//@description: 输入币种及价格返回人民币价格
//@param: price string
//@return: cnyPrice float64, err error

func (r *RateConverter) GetCNY(price string) (cnyPrice float64, err error) {
	currencySignOrName, currencyPrice, err := utils.MatchCurrencyAndPrice(price) //分别提取输入字符串中的币种缩写/符号、价格
	if err != nil {
		return
	}
	currencyName, err := r.getCurrencyName(currencySignOrName) //匹配币种名称
	if err != nil {
		return
	}
	converterContext := &ConverterContext{}
	switch currencyName { //判断对应币种换算策略
	case "USD":
		converterContext.Strategy = &strategy.USDstrategy{
			Price: currencyPrice,
		}
	case "GBP":
		converterContext.Strategy = &strategy.UBPstrategy{
			Price: currencyPrice,
		}
	case "EUR":
		converterContext.Strategy = &strategy.EURstrategy{
			Price: currencyPrice,
		}
	case "HKD":
		converterContext.Strategy = &strategy.HKDstrategy{
			Price: currencyPrice,
		}
	case "JPY":
		converterContext.Strategy = &strategy.JPYstrategy{
			Price: currencyPrice,
		}
	}
	cnyPrice, err = converterContext.ConverterToCNY()
	return
}

//@author: [Gaayean]
//@function: getCurrencyName
//@description: 根据提取的货币字符串返回货币的英文缩写，如果提取的货币不在系统支持范围内则返回错误消息
//@param: sign string
//@return: currencyName string, err error

func (r *RateConverter) getCurrencyName(sign string) (currencyName string, err error) {
	if len(global.RATE_CONFIG.Currencies) == 0 {
		err = fmt.Errorf("无可转换货币，请检查系统配置文件或检查是否有读取配置文件")
		return
	}
	for _, v := range global.RATE_CONFIG.Currencies { //
		if sign == v.Name || sign == v.Sign {
			currencyName = v.Name
			return
		}
	}
	err = fmt.Errorf("暂不支持此货币转换")
	return
}
