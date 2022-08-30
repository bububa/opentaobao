package ascp

// TmsOrderConfirmResponse 结构体
type TmsOrderConfirmResponse struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 业务请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否业务异常
	BizException bool `json:"biz_exception,omitempty" xml:"biz_exception,omitempty"`
}
