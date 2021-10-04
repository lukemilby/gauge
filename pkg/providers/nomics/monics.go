package nomics

import "time"

type CurrencyTracker struct {
	Currency             string    `json:"currency"`
	ID                   string    `json:"id"`
	Status               string    `json:"status"`
	Price                string    `json:"price"`
	PriceDate            time.Time `json:"price_date"`
	PriceTimestamp       time.Time `json:"price_timestamp"`
	Symbol               string    `json:"symbol"`
	CirculatingSupply    string    `json:"circulating_supply"`
	MaxSupply            string    `json:"max_supply"`
	Name                 string    `json:"name"`
	LogoURL              string    `json:"logo_url"`
	MarketCap            string    `json:"market_cap"`
	MarketCapDominance   string    `json:"market_cap_dominance"`
	TransparentMarketCap string    `json:"transparent_market_cap"`
	NumExchanges         string    `json:"num_exchanges"`
	NumPairs             string    `json:"num_pairs"`
	NumPairsUnmapped     string    `json:"num_pairs_unmapped"`
	FirstCandle          time.Time `json:"first_candle"`
	FirstTrade           time.Time `json:"first_trade"`
	FirstOrderBook       time.Time `json:"first_order_book"`
	FirstPricedAt        time.Time `json:"first_priced_at"`
	Rank                 string    `json:"rank"`
	RankDelta            string    `json:"rank_delta"`
	High                 string    `json:"high"`
	HighTimestamp        string    `json:"high_timestamp"`
	OneD                 struct {
		PriceChange                   string `json:"price_change"`
		PriceChangePct                string `json:"price_change_pct"`
		Volume                        string `json:"volume"`
		VolumeChange                  string `json:"volume_change"`
		VolumeChangePct               string `json:"volume_change_pct"`
		MarketCapChange               string `json:"market_cap_change"`
		MarketCapChangePct            string `json:"market_cap_change_pct"`
		TransparentMarketCapChange    string `json:"transparent_market_cap_change"`
		TransparentMarketCapChangePct string `json:"transparent_market_cap_change_pct"`
		VolumeTransparency            []struct {
			Grade           string `json:"grade"`
			Volume          string `json:"volume"`
			VolumeChange    string `json:"volume_change"`
			VolumeChangePct string `json:"volume_change_pct"`
		} `json:"volume_transparency"`
	} `json:"1d"`
}