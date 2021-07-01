package aliexpresssumaitong

// Lines 结构体
type Lines struct {
	// 商品数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 商品未含税总价
	OrderTaxExcludedAmount string `json:"order_tax_excluded_amount,omitempty" xml:"order_tax_excluded_amount,omitempty"`
	// 商品运费总价
	DeliveryAmount string `json:"delivery_amount,omitempty" xml:"delivery_amount,omitempty"`
	// 商品税金总价
	OrderTaxAmount string `json:"order_tax_amount,omitempty" xml:"order_tax_amount,omitempty"`
	// 运费税金总价
	DeliveryTaxAmount string `json:"delivery_tax_amount,omitempty" xml:"delivery_tax_amount,omitempty"`
	// 商品税率
	TaxRates []string `json:"tax_rates,omitempty" xml:"tax_rates>string,omitempty"`
	// 运费税率
	DeliveryTaxRates []string `json:"delivery_tax_rates,omitempty" xml:"delivery_tax_rates>string,omitempty"`
	// 子订单号id
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
}
