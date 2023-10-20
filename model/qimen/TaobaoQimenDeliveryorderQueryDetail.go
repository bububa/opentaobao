package qimen

// TaobaoQimenDeliveryorderQueryDetail 结构体
type TaobaoQimenDeliveryorderQueryDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
