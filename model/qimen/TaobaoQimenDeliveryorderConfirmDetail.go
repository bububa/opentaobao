package qimen

// TaobaoQimenDeliveryorderConfirmDetail 结构体
type TaobaoQimenDeliveryorderConfirmDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
