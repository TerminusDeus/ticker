package main

import (
	"fmt"
	"github.com/toorop/go-bittrex"
	"os"
	"strings"
)

func Buy(marketName string) {
	bittrex := bittrex.New(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	coin := strings.Split(marketName, "-")[0]
	bal, _ := bittrex.GetBalance(coin)
	if bal.Available < 0.0001 {
		fmt.Println("Not enough captial to buy")
		return
	}
}

func Sell(marketName string) {
	bittrex := bittrex.New(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	coin := strings.Split(marketName, "-")[1]
	bal, _ := bittrex.GetBalance(coin)
	if bal.Available < 0.0001 {
		fmt.Println("Nothing to sell")
		return
	}
	uuid, err := bittrex.SellMarket(marketName, bal.Available)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Sale completed", uuid)
}
