package alitripmerchant

// OrderCountDetailVO 结构体
type OrderCountDetailVO struct {
	// 状态枚举 code
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 状态中文值
	OrderStatusStr string `json:"order_status_str,omitempty" xml:"order_status_str,omitempty"`
	// 对应数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}
