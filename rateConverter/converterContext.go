package rateconverter

//汇率换算上下文结构体
//@author: [Gaayean]
type ConverterContext struct {
	Strategy RateStrategy
}

// func NewConverterContext(rateStrategy RateStrategy) *ConverterContext {
// 	return &ConverterContext{
// 		Strategy: rateStrategy,
// 	}
// }

func (c *ConverterContext) ConverterToCNY() (float64, error) {
	return c.Strategy.ConverterToCNY()
}
