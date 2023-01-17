package wdk

// FulfillProduct 结构体
type FulfillProduct struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品履约线路
	LineInstances string `json:"line_instances,omitempty" xml:"line_instances,omitempty"`
}
