package alitripmerchant

// AlitripmerchantgalaxyderbymembervouchercardordercancelResponse 结构体
type AlitripmerchantgalaxyderbymembervouchercardordercancelResponse struct {
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 订单取消成功
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
