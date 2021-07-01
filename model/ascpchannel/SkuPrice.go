package ascpchannel

// SkuPrice 结构体
type SkuPrice struct {
	// 扩展价格
	ExtendPrice string `json:"extend_price,omitempty" xml:"extend_price,omitempty"`
	// 价格值
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 币种
	CurrencyPriceValue string `json:"currency_price_value,omitempty" xml:"currency_price_value,omitempty"`
}
