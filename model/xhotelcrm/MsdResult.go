package xhotelcrm

// MsdResult 结构体
type MsdResult struct {
	// 系统异常
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// S_SYS_EXCEPTION
	ErrCode bool `json:"err_code,omitempty" xml:"err_code,omitempty"`
}
