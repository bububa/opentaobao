package qimen

// TaobaoqimendeliveryordercreateDetail 结构体
type TaobaoqimendeliveryordercreateDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
