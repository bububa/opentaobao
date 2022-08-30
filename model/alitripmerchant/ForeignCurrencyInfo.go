package alitripmerchant

// ForeignCurrencyInfo 结构体
type ForeignCurrencyInfo struct {
	// 外币每日价格
	DailyPriceList []DailyPrice `json:"daily_price_list,omitempty" xml:"daily_price_list>daily_price,omitempty"`
	// 外币支付金额，含税
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 外币基础价格,不包含税费等
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 外币总税价
	TotalTax string `json:"total_tax,omitempty" xml:"total_tax,omitempty"`
	// 外币币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
}
