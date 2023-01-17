package wdk

// OrderQueryResult 结构体
type OrderQueryResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 订单数据
	Result *OrderSuccessResponse `json:"result,omitempty" xml:"result,omitempty"`
	// 接口状态
	State bool `json:"state,omitempty" xml:"state,omitempty"`
}
