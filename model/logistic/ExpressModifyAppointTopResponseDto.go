package logistic

// ExpressModifyAppointTopResponseDto 结构体
type ExpressModifyAppointTopResponseDto struct {
	// 订单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
