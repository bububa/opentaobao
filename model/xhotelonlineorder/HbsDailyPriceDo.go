package xhotelonlineorder

// HbsDailyPriceDo 结构体
type HbsDailyPriceDo struct {
	// 每日价格
	DailyPrices []DailyPriceTo `json:"daily_prices,omitempty" xml:"daily_prices>daily_price_to,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 汇率
	ExchangeRate string `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
}
