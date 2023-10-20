package qimen

// TaobaoQimenDeliveryorderBatchconfirmDetail 结构体
type TaobaoQimenDeliveryorderBatchconfirmDetail struct {
	// 订单商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
