package legalcase

// ServiceResult 结构体
type ServiceResult struct {
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// errorcode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 内容
	Contents []Content `json:"contents,omitempty" xml:"contents>content,omitempty"`
	// errormasg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// content
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 错误编码
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
}
