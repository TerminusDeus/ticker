package main

import (
        "fmt"
        "github.com/toorop/go-bittrex"
        "time"
        "sort"
        "math"
)

const (
        API_KEY            = os.Getenv("API_KEY")
        API_SECRET         = os.Getenv("API_SECRET")
        TICKER_SAMPLE_SIZE = 5
        HISTORY_LENGTH     = 50
)

type Market struct {
        Name    string
        Avg     float64
        Last    float64
        History []float64
}

func NewMarket(name string, interval time.Duration) *Market {
        var noHistory = make([]float64, HISTORY_LENGTH)

        m := Market{name, 0, 0, noHistory}
        go m.Ticker(interval)
        m.getMarketHistory()

        return &m
}

func (m *Market) getMarketHistory() {
        bittrex := bittrex.New(API_KEY, API_SECRET)
        ticks, err := bittrex.GetTicks(m.Name, "hour")
        if err != nil {
                fmt.Println(err)
        }
        reverse(ticks)
        for i, t := range(ticks) {
                if i >= 50 {
                        break
                }
                m.History[i] = t.Close
        }
}

func reverse(numbers []bittrex.Candle) {
        for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
                numbers[i], numbers[j] = numbers[j], numbers[i]
        }
}

func (m *Market) calcLast() {
        bittrex := bittrex.New(API_KEY, API_SECRET)
        var sampleSet = make([]float64, TICKER_SAMPLE_SIZE)
        i := 0
        for i < TICKER_SAMPLE_SIZE {
                ticker, err := bittrex.GetTicker(m.Name)
                if err != nil {
                        fmt.Println(err)
                }
                sampleSet[i] = ticker.Last
                time.Sleep(1 * time.Second)
                i += 1
        }
        sort.Float64s(sampleSet)
        m.Last = sampleSet[int(math.Floor(TICKER_SAMPLE_SIZE / 2 + .5))]
}

func (m *Market) Ticker(d time.Duration) {
        for {
                m.calcLast()
                m.History = append(m.History, m.Last)[1:]
                var total float64 = 0
                for _, v := range m.History {
                        total += v
                }
                m.Avg = total / float64(len(m.History))
                time.Sleep(d)
        }
}
