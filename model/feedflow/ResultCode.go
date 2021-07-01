package feedflow

// ResultCode 结构体
type ResultCode struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
