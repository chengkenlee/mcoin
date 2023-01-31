package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Gateios struct {
	Result string `json:"result"`
	Data   []struct {
		No          int         `json:"no"`
		Symbol      string      `json:"symbol"`
		Name        string      `json:"name"`
		NameEn      string      `json:"name_en"`
		NameCn      string      `json:"name_cn"`
		Pair        string      `json:"pair"`
		Rate        string      `json:"rate"`
		VolA        string      `json:"vol_a"`
		VolB        string      `json:"vol_b"`
		CurrA       string      `json:"curr_a"`
		CurrB       string      `json:"curr_b"`
		CurrSuffix  string      `json:"curr_suffix"`
		RatePercent string      `json:"rate_percent"`
		Trend       string      `json:"trend"`
		Supply      interface{} `json:"supply"`
		Marketcap   interface{} `json:"marketcap"`
		Lq          string      `json:"lq"`
		PRate       int         `json:"p_rate"`
		High        string      `json:"high"`
		Low         string      `json:"low"`
	} `json:"data"`
}

//格式化日志
func Gets() string {
	client := &http.Client{}
	//读取配置文件的url, 并提交请求
	resquest, err := http.NewRequest("GET", "https://data.gateapi.io/api2/1/marketlist", nil)
	if err != nil {
		println(fmt.Sprintf("%s", err))
	}
	//处理返回的结果
	response, err0 := client.Do(resquest)
	if err0 != nil {
		println(fmt.Sprintf("%s", err0))
	}
	//检出结果集
	body, err1 := ioutil.ReadAll(response.Body)
	//关闭流
	defer response.Body.Close()
	if err1 != nil {
		println(fmt.Sprintf("%s", err1))
	}
	response.Body.Close()
	if err != nil {
		println(fmt.Sprintf("%s", err))
	}
	return string(body)
}

func main() {
	var coin string
	flag.StringVar(&coin, "k", "", "<btc/eth/etc>")
	flag.Parse()
	if len(os.Args) < 3 {
		flag.PrintDefaults()
	} else {
		//获取json
		var rs Gateios
		jsons := Gets()
		err := json.Unmarshal([]byte(jsons), &rs)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		println(fmt.Sprintf("%-8s%-18s%-15s%-15s%-18s%-13s%-13s%-13s%-10s%-15s%-18s%-5s%-5s",
			"Symbol", "NameEn", "Pair", "Rate", "VolA", "VolB", "CurrSuffix", "RatePercent", "Trend", "Supply", "Marketcap", "Lq", "NameCn"))
		for _, item := range rs.Data {
			if strings.ToUpper(coin) == item.Symbol {
				fmt.Println(fmt.Sprintf("%-8s%-18s%-15s%-15s%-18s%-13s%-13s%-13s%-10s%-15s%-18s%-5s%-5s",
					item.Symbol, item.NameEn, item.Pair, item.Rate, item.VolA, item.VolB, item.CurrSuffix, item.RatePercent, item.Trend, item.Supply, item.Marketcap, item.Lq, item.NameCn))
			}
		}
	}
}
