package yunos

// DpResult 结构体
type DpResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
