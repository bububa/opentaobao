package aliexpresssumaitong

// PreCreateOrderRequest 结构体
type PreCreateOrderRequest struct {
	// 商品信息
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}
