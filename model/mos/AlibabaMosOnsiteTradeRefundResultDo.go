package mos

// AlibabamosonsitetraderefundResultDo 结构体
type AlibabamosonsitetraderefundResultDo struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *RefundResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
