package ascpchannel

// Additionalinfodto 结构体
type Additionalinfodto struct {
	// 未税价格 (精度 6位) 采购&amp;退供 经销必填
	PriceNoTax string `json:"price_no_tax,omitempty" xml:"price_no_tax,omitempty"`
}
