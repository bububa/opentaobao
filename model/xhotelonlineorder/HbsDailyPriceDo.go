package xhotelonlineorder

import (
	"sync"
)

// HbsDailyPriceDo 结构体
type HbsDailyPriceDo struct {
	// 每日价格
	DailyPrices []DailyPriceTo `json:"daily_prices,omitempty" xml:"daily_prices>daily_price_to,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 汇率
	ExchangeRate string `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
}

var poolHbsDailyPriceDo = sync.Pool{
	New: func() any {
		return new(HbsDailyPriceDo)
	},
}

// GetHbsDailyPriceDo() 从对象池中获取HbsDailyPriceDo
func GetHbsDailyPriceDo() *HbsDailyPriceDo {
	return poolHbsDailyPriceDo.Get().(*HbsDailyPriceDo)
}

// ReleaseHbsDailyPriceDo 释放HbsDailyPriceDo
func ReleaseHbsDailyPriceDo(v *HbsDailyPriceDo) {
	v.DailyPrices = v.DailyPrices[:0]
	v.CurrencyCode = ""
	v.ExchangeRate = ""
	poolHbsDailyPriceDo.Put(v)
}
