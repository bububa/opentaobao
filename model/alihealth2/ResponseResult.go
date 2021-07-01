package alihealth2

// ResponseResult 结构体
type ResponseResult struct {
	// 订单详情对象
	Result *TopOrderDetailDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
}
