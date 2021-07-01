package perfect

// PerfectSalePropertyDto 结构体
type PerfectSalePropertyDto struct {
	// 属性key
	SalePropertyKey string `json:"sale_property_key,omitempty" xml:"sale_property_key,omitempty"`
	// 自定义属性值
	SalePropertyValue string `json:"sale_property_value,omitempty" xml:"sale_property_value,omitempty"`
}
