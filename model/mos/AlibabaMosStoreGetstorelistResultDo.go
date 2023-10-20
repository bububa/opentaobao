package mos

// AlibabamosstoregetstorelistResultDo 结构体
type AlibabamosstoregetstorelistResultDo struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// data
	Data *MjStoresTopVo `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
