package main

import (
        "fmt"
        "github.com/d4l3k/talib"
        "time"
        "flag"
)

func main() {

        strategy := *flag.String("strategy", "sma", "The trading strategy")
        //interval := *flag.Int("interval", 5, "The refresh frequency")

        flag.Parse()

        markets := make([]*Market, 0)
        for _, marketName := range(flag.Args()) {
                markets = append(markets, NewMarket(marketName, 3 * time.Second))
        }

        fmt.Println("Using strategy: ", strategy, "At intervals: "/*, interval*/)

        for {
                for _, m := range(markets) {
                        time.Sleep(5 * time.Second)
                        sma := talib.Sma(m.History, HISTORY_LENGTH)[0]
                        fmt.Println("SMA:  ", sma)
                        ema := talib.Ema(m.History, HISTORY_LENGTH)[0]

                        fmt.Println(m.Name)
                        fmt.Println("SMA:  ", sma)
                        fmt.Println("EMA:  ", ema)
                        fmt.Println("LAST: ", m.Last)

                        if (m.Last > sma) {
                                fmt.Println("BUY")
				Buy(m.Name)
                        } else {
                                fmt.Println("SELL")
				Sell(m.Name)
                        }
                        fmt.Println()
                }
        }
}
