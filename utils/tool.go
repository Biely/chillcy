package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//@author: [Gaayean]
//@function: MatchCurrencyAndPrice
//@description: 从输入字符串中提取货币缩写或符号、金额
//@param: inputPrice string
//@return: sign string, price float64, err error

func MatchCurrencyAndPrice(inputPrice string) (sign string, price float64, err error) {
	if len(inputPrice) == 0 {
		err = fmt.Errorf("输入内容为空")
		return
	}
	noDotPrice := strings.ReplaceAll(inputPrice, ",", "") //去除字符串中的千位逗号
	priceRegex := `\d+(\.\d+)?`
	rePrice, err := regexp.Compile(priceRegex)
	if err != nil {
		return
	}
	reSign, err := regexp.Compile(fmt.Sprintf("^[^%v]+", priceRegex)) //这里是匹配除了价格之外的字符串,可能不够严谨，可以选择改为（\$|£|€|HK$）这样的写法，不过会越来越长
	if err != nil {
		return
	}
	sign = reSign.FindString(noDotPrice)
	if sign == "" {
		err = fmt.Errorf("未知货币类型")
		return
	}
	priceStr := rePrice.FindString(noDotPrice)
	if priceStr == "" {
		err = fmt.Errorf("未知货币价格")
		return
	}
	price, err = strconv.ParseFloat(priceStr, 64)
	return
}
