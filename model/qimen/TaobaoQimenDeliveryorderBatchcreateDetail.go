package qimen

// TaobaoQimenDeliveryorderBatchcreateDetail 结构体
type TaobaoQimenDeliveryorderBatchcreateDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
