package product

// SkuStatus 结构体
type SkuStatus struct {
	// sku状态描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// sku商家编码和sku_id两者取一
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// outerId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// sku状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
