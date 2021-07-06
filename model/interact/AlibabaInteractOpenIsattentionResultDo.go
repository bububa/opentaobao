package interact

// AlibabaInteractOpenIsattentionResultDo 结构体
type AlibabaInteractOpenIsattentionResultDo struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// data
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// isRetry
	IsRetry bool `json:"is_retry,omitempty" xml:"is_retry,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
