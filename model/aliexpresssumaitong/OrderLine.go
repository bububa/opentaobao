package aliexpresssumaitong

// OrderLine 结构体
type OrderLine struct {
	// 子订单扩展字段
	ExtraParams string `json:"extra_params,omitempty" xml:"extra_params,omitempty"`
	// 子订单号
	OrderLineId int64 `json:"order_line_id,omitempty" xml:"order_line_id,omitempty"`
}
