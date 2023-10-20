package wdk

// AlibabawdkordersharestockcpsorderlistApiResult 结构体
type AlibabawdkordersharestockcpsorderlistApiResult struct {
	// 调用接口返回对象
	Model []CpsOrderResponse `json:"model,omitempty" xml:"model>cps_order_response,omitempty"`
	// 调用接口返回错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用接口返回错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用接口成功失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
