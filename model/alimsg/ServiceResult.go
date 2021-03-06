package alimsg

// ServiceResult 结构体
type ServiceResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 内容
	Content *SendTemplateMsgResponse `json:"content,omitempty" xml:"content,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
