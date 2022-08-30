package wdk

// SubOrderReturn 结构体
type SubOrderReturn struct {
	// 拦截状态
	InterceptStatus string `json:"intercept_status,omitempty" xml:"intercept_status,omitempty"`
	// 子订单
	SubOrderCode string `json:"sub_order_code,omitempty" xml:"sub_order_code,omitempty"`
	// 取消结果
	CancelSuccess bool `json:"cancel_success,omitempty" xml:"cancel_success,omitempty"`
}
