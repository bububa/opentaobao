package util

// ResultCode 结构体
type ResultCode struct {
	// 错误码描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误妈code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
