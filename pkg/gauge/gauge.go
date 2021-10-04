package gauge

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/lukemilby/gauge/pkg/providers/nomics"
)

var endpoint = "https://api.nomics.com/v1"

type Gauge struct {
	config Config
}

type Config struct {
	Interval string // this sets the timeframe for the data ex. 1m 1M 3D 7h
	ApiKey   string
}

func NewGauge(config Config) *Gauge {
	g := &Gauge{
		config,
	}

	return g
}

//TODO parse
//TODO list currencys
func (g *Gauge) Run() error {

	url := "/currencies/ticker"

	stableCoins := "USDC,USDT,BUSD,DAI"

	var payload []nomics.CurrencyTracker

	r := resty.New()

	resp, err := r.R().
		SetQueryParams(map[string]string{
			"key":      g.config.ApiKey,
			"ids":      stableCoins,
			"interval": "1d",
			"convert":  "USD",
			"per-page": "10",
			"page":     "1",
		}).
		EnableTrace().
		Get(endpoint + url)
	if err != nil {
		return errors.New(
			fmt.Sprintf("error: %s\n", err.Error()),
		)
	}

	err = json.Unmarshal(resp.Body(), &payload)
	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range payload {
		fmt.Println(item.Name)
	}

	return nil
}
