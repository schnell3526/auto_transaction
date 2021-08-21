package gmo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/schnell3526/auto_auto_transaction/src/gmo/symbol"
)

// 取引所の稼働状態を取得
func GetServiceStatus(JTC bool) *ServiceStatus {
	endPoint := "https://api.coin.z.com/public"
	path := "/v1/status"

	response, _ := http.Get(endPoint + path)
	body, _ := ioutil.ReadAll(response.Body)

	var result ServiceStatus
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if JTC {
		tokyo, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		result.Responsetime = result.Responsetime.In(tokyo)
	}

	return &result
}

// 最新レートを取得
func GetLatestRate(coin_type symbol.SymbolType, JTC bool) *LatestRate {
	endPoint := "https://api.coin.z.com/public"
	path := "/v1/ticker"

	coin := symbol.NewSymbol(coin_type)
	path += "?symbol=" + coin.Literal

	response, _ := http.Get(endPoint + path)
	body, _ := ioutil.ReadAll(response.Body)

	var result LatestRate
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if JTC {
		tokyo, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		result.Data[0].Timestamp = result.Data[0].Timestamp.In(tokyo)
		result.Responsetime = result.Responsetime.In(tokyo)
	}

	return &result
}

// 板情報を取得
func GetOrderBook(coin_type symbol.SymbolType, JTC bool) (*OrderBook, error) {
	endPoint := "https://api.coin.z.com/public"
	path := "/v1/orderbooks"

	coin := symbol.NewSymbol(coin_type)
	path += "?symbol=" + coin.Literal

	response, _ := http.Get(endPoint + path)
	body, _ := ioutil.ReadAll(response.Body)

	var result OrderBook
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if JTC {
		tokyo, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		result.Responsetime = result.Responsetime.In(tokyo)
	}

	return &result, nil
}

// // TODO:取引履歴を取得
// func GetTradeHistory()

// // TODO:KLine情報を取得
// func GetKLineData()
