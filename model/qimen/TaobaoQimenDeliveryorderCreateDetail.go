package qimen

// TaobaoQimenDeliveryorderCreateDetail 
type TaobaoQimenDeliveryorderCreateDetail struct {
    // 商品列表
    Items   []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
