package tmallservice

// SimpleServiceSkuDtOs 结构体
type SimpleServiceSkuDtOs struct {
	// 服务skuCode
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 服务sku名称
	DisplayName string `json:"display_name,omitempty" xml:"display_name,omitempty"`
}
