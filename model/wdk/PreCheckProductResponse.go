package wdk

// PreCheckProductResponse 结构体
type PreCheckProductResponse struct {
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 是否可作业
	CanFulfill bool `json:"can_fulfill,omitempty" xml:"can_fulfill,omitempty"`
}
