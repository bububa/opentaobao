package aliexpresssumaitong

// Order 结构体
type Order struct {
	// 子订单信息
	OrderLines []OrderLine `json:"order_lines,omitempty" xml:"order_lines>order_line,omitempty"`
	// 外部订单号
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 订单扩展字段
	ExtraParams *ExtraMap `json:"extra_params,omitempty" xml:"extra_params,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}
