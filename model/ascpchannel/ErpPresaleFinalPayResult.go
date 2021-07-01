package ascpchannel

// ErpPresaleFinalPayResult 结构体
type ErpPresaleFinalPayResult struct {
	// 服务调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
}
