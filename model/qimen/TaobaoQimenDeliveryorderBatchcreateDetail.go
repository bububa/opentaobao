package qimen

// TaobaoQimenDeliveryorderBatchcreateDetail 
type TaobaoQimenDeliveryorderBatchcreateDetail struct {
    // 商品列表
    Items   []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
