package alisports

// AlispResult 结构体
type AlispResult struct {
	// 错误信息
	AlispMsg string `json:"alisp_msg,omitempty" xml:"alisp_msg,omitempty"`
	// 数据对象
	AlispData *AuthAccountInfoDto `json:"alisp_data,omitempty" xml:"alisp_data,omitempty"`
	// 错误码
	AlispCode int64 `json:"alisp_code,omitempty" xml:"alisp_code,omitempty"`
}
