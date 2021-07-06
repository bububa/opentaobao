package nlife

// Deliveritemlist 结构体
type Deliveritemlist struct {
	// 商品唯一码列表
	UniqueCodeList []string `json:"unique_code_list,omitempty" xml:"unique_code_list>string,omitempty"`
	// 商品条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 发货的数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
