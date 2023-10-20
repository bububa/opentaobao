package qimen

// TaobaoqimendeliveryorderqueryDetail 结构体
type TaobaoqimendeliveryorderqueryDetail struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
