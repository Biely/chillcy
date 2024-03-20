package main

import (
	"fmt"

	"chillcy.com/rate-converter/initialize"
	rateconverter "chillcy.com/rate-converter/rateConverter"
)

//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	initialize.Viper()                                //初始化viper
	converterFunc := new(rateconverter.RateConverter) //实例化汇率转换对象
	price, err := converterFunc.GetCNY("€499.99")     //调用转人民币方法
	fmt.Println(price, err)
}
