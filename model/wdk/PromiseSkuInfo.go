package wdk

// PromiseSkuInfo 结构体
type PromiseSkuInfo struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品履约线路
	LineInstances string `json:"line_instances,omitempty" xml:"line_instances,omitempty"`
	// 加购数量
	Quantity *BigDecimal `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
