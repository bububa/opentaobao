package idleisv

// AppraiseIsvDistributionDto 结构体
type AppraiseIsvDistributionDto struct {
	// 分销商利润（单位分），global_tax_fee + product_price + daigou_fee 等于订单总价
	DaigouFee string `json:"daigou_fee,omitempty" xml:"daigou_fee,omitempty"`
	// 商品的供货价格，即货品的价格（单位分）
	ProductPrice string `json:"product_price,omitempty" xml:"product_price,omitempty"`
	// 国际商品税费（单位分）
	GlobalTaxFee string `json:"global_tax_fee,omitempty" xml:"global_tax_fee,omitempty"`
}
