package train

// TapResult 结构体
type TapResult struct {
	// 返回code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回msg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 锁单rs
	Module *LockOrderRs `json:"module,omitempty" xml:"module,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
