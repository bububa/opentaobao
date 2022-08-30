package product

// ItemSkuStatus 结构体
type ItemSkuStatus struct {
	// sku集合
	SkuStatusList []SkuStatus `json:"sku_status_list,omitempty" xml:"sku_status_list>sku_status,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
