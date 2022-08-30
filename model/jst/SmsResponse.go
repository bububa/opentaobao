package jst

// SmsResponse 结构体
type SmsResponse struct {
	// 请求成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 请求成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Module *SmsTaskModel `json:"module,omitempty" xml:"module,omitempty"`
	// 请求成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
