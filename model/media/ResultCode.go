package media

// ResultCode 结构体
type ResultCode struct {
	// 错误描述
	Info string `json:"info,omitempty" xml:"info,omitempty"`
	// 错误代码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
}
