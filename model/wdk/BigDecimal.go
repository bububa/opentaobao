package wdk

// BigDecimal 结构体
type BigDecimal struct {
	// 每份重量
	SkuWeight *BigDecimal `json:"sku_weight,omitempty" xml:"sku_weight,omitempty"`
}
