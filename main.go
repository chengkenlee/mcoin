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
	Result string `bson:"result"`
	Data   []struct {
		No           string `bson:"no"`
		Symbol       string `bson:"symbol"`
		Name         string `bson:"name"`
		Name_en      string `bson:"name_en"`
		Name_cn      string `bson:"name_cn"`
		Pair         string `bson:"pair"`
		Rate         string `bson:"rate"`
		Vol_a        string `bson:"vol_a"`
		Vol_b        string `bson:"vol_b"`
		Curr_a       string `bson:"curr_a"`
		Curr_b       string `bson:"curr_b"`
		Curr_suffix  string `bson:"curr_suffix"`
		Rate_percent string `bson:"rate_percent"`
		Trend        string `bson:"trend"`
		Supply       string `bson:"supply"`
		Marketcap    string `bson:"marketcap"`
		Lq           string `bson:"lq"`
		P_rate       string `bson:"p_rate"`
	} `bson:"data"`
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
		json.Unmarshal([]byte(jsons), &rs)
		println(fmt.Sprintf("%-8s%-10s%-10s%-15s%-18s%-13s%-13s%-13s%-10s%-15s%-18s%-5s%-5s", "SYMBOL", "NAME_EN", "PAIR", "RATE", "VOL_A", "VOL_B", "CURR_SUFFIX", "RATE_PERCENT", "TREND", "SUPPLY", "MARKETCAP", "LQ", "NAME_CN"))
		for _, item := range rs.Data {
			if strings.ToUpper(coin) == item.Symbol {
				println(fmt.Sprintf("%-8s%-10s%-10s%-15s%-18s%-13s%-13s%-13s%-10s%-15s%-18s%-5s%-5s", item.Symbol, item.Name_en, item.Pair, item.Rate, item.Vol_a, item.Vol_b, item.Curr_suffix, item.Rate_percent, item.Trend, item.Supply, item.Marketcap, item.Lq, item.Name_cn))
			}
		}
	}
}
