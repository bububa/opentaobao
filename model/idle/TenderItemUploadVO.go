package idle

// TenderItemUploadVO 结构体
type TenderItemUploadVO struct {
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 类目
	Category string `json:"category,omitempty" xml:"category,omitempty"`
}
