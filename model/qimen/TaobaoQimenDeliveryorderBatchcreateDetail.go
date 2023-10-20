package qimen

// TaobaoqimendeliveryorderbatchcreateDetail 结构体
type TaobaoqimendeliveryorderbatchcreateDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
