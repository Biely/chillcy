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
	priceRegex := `\d+(.\d+)?`
	rePrice, err := regexp.Compile(priceRegex)
	if err != nil {
		return
	}
	reSign, err := regexp.Compile(fmt.Sprintf("[^%v]+", priceRegex))
	if err != nil {
		return
	}
	sign = reSign.FindString(noDotPrice)
	priceStr := rePrice.FindString(noDotPrice)
	price, err = strconv.ParseFloat(priceStr, 64)
	return
}
