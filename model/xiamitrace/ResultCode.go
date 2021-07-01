package xiamitrace

// ResultCode 结构体
type ResultCode struct {
	// result message
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// result code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
