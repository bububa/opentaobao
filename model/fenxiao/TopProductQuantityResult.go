package fenxiao

// TopProductQuantityResult 结构体
type TopProductQuantityResult struct {
	// 更新后库存数量
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 更新时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 产品数字ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// SKU ID
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
