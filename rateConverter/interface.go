package rateconverter

//定义汇率换算策略接口
//@author: [Gaayean]
type RateStrategy interface {
	ConverterToCNY() (float64, error)
}
