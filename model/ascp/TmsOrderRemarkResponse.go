package ascp

// TmsOrderRemarkResponse 结构体
type TmsOrderRemarkResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 业务错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 业务返回数据
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
