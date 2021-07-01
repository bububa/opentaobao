package qimen

// OrderLines 结构体
type OrderLines struct {
	// 订单详情
	OrderLine *OrderLine `json:"orderLine,omitempty" xml:"orderLine,omitempty"`
}
