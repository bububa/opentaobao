package xiamiatrist

// ResultCode 结构体
type ResultCode struct {
	// 是否成功
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
