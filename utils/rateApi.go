package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RateData struct {
	Base       string             `json:"base"`
	LastDataAt int64              `json:"last_data_at"`
	Rates      map[string]float64 `json:"rates"`
}

//@author: [Gaayean]
//@function: GetRateToCNY
//@description: 通过接口获取目标货币转人民币的汇率
//@param: url string, currencyName string
//@return: cnyRate float64, err error

func GetRateToCNY(url string, currencyName string) (cnyRate float64, err error) {
	resRateData, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("请求汇率失败：%v", err)
		return
	}
	defer resRateData.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resRateData.Body)
	if err != nil {
		err = fmt.Errorf("读取响应体失败: %v", err)
		return
	}
	var rateData RateData
	err = json.Unmarshal(body, &rateData)
	if err != nil {
		err = fmt.Errorf("解析 JSON 失败: %v", err)
		return
	}
	rateToCNY := rateData.Rates["CNY"]
	rateToInputCurr := rateData.Rates[currencyName]
	cnyRate = float64(1) / rateToInputCurr * rateToCNY
	// fmt.Println(cnyRate)
	return
}
