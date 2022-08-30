package alihealthpw

// ResponseMessage 结构体
type ResponseMessage struct {
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 数据
	Data *ApplyDetailResp `json:"data,omitempty" xml:"data,omitempty"`
}
