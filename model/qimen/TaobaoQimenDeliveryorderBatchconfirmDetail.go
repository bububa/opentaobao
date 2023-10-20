package qimen

// TaobaoqimendeliveryorderbatchconfirmDetail 结构体
type TaobaoqimendeliveryorderbatchconfirmDetail struct {
	// 订单商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
