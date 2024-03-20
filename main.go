package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Biely/chillcy/initialize"
	rateconverter "github.com/Biely/chillcy/rateConverter"
)

//模拟测试输入

//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	initialize.Viper()                                //初始化viper
	converterFunc := new(rateconverter.RateConverter) //实例化汇率转换对象
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("欢迎使用汇率转换工具")
	fmt.Println("请输入“币种+金额”例如：$99.99，以回车键结束（输入exit退出程序）")
	for scanner.Scan() {
		inputPrice := scanner.Text()
		if inputPrice == "exit" {
			break
		}
		fmt.Println("您输入的是：", inputPrice)
		price, err := converterFunc.GetCNY(inputPrice) //调用转人民币方法
		if err != nil {
			fmt.Println("转换失败，失败原因：", err)
		}
		fmt.Println("转换人民币结果：", price)
		fmt.Println("请继续输入（输入exit退出程序）")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "读取错误:", err)
	}
	fmt.Println("程序已退出")
}
