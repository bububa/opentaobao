package qimen

// TaobaoqimendeliveryorderconfirmDetail 结构体
type TaobaoqimendeliveryorderconfirmDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
