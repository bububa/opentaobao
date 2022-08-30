package qimen

// RelatedOrders 结构体
type RelatedOrders struct {
	// 关联的订单类型,关联的订单类型,string(50),,
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 关联的订单编号,关联的订单编号,string(50),,
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
}
