package mos

// AlibabaMjOcSyncpayinfoResultDo 结构体
type AlibabaMjOcSyncpayinfoResultDo struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Date string `json:"date,omitempty" xml:"date,omitempty"`
	// 错误码
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
