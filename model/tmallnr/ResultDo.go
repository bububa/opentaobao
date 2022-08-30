package tmallnr

// ResultDo 结构体
type ResultDo struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// DTO
	Data *NrtCrmSellerInfoDto `json:"data,omitempty" xml:"data,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
