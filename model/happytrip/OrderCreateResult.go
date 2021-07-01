package happytrip

// OrderCreateResult 结构体
type OrderCreateResult struct {
	// 订单信息
	Order *OrderInfo `json:"order,omitempty" xml:"order,omitempty"`
	// 价格信息
	Price *PriceInfo `json:"price,omitempty" xml:"price,omitempty"`
}
