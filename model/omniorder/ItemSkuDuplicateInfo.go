package omniorder

// ItemSkuDuplicateInfo 结构体
type ItemSkuDuplicateInfo struct {
	// 重复商品详情列表，如果重复商品过多，目前只展示部分
	DuplicateDetails []ItemSkuDuplicateDetail `json:"duplicate_details,omitempty" xml:"duplicate_details>item_sku_duplicate_detail,omitempty"`
	// barcode
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// outerId
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// skuBarcode
	SkuBarcode string `json:"sku_barcode,omitempty" xml:"sku_barcode,omitempty"`
	// skuOuterId
	SkuOuterId string `json:"sku_outer_id,omitempty" xml:"sku_outer_id,omitempty"`
	// 重复商品数量
	DuplicateSize int64 `json:"duplicate_size,omitempty" xml:"duplicate_size,omitempty"`
}
