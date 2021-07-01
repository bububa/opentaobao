package qimen

// TaobaoQimenDeliveryorderCreateDetail 结构体
type TaobaoQimenDeliveryorderCreateDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
