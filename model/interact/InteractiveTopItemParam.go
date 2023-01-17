package interact

// InteractiveTopItemParam 结构体
type InteractiveTopItemParam struct {
	// url中的自定义参数
	OutKey string `json:"out_key,omitempty" xml:"out_key,omitempty"`
	// 商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// skuId
	Sku int64 `json:"sku,omitempty" xml:"sku,omitempty"`
	// 购买数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
