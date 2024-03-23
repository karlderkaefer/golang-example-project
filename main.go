package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var version = "development"

type Crypto struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type Price struct {
	USD float64 `json:"usd"`
}

type MarketData struct {
	CurrentPrice Price `json:"current_price"`
}

type CryptoData struct {
	ID         string     `json:"id"`
	Symbol     string     `json:"symbol"`
	Name       string     `json:"name"`
	MarketData MarketData `json:"market_data"`
}

func printVersion() string {
	version := fmt.Sprintf("Version: %s", version)
	fmt.Println(version)
	return version
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		printVersion()
		return
	}

	cryptoID := "ethereum"

	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/" + cryptoID)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var data CryptoData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("The current price of %s (%s) is $%f\n", data.Name, data.Symbol, data.MarketData.CurrentPrice.USD)
}
