package wdk

// AlibabatxcsbrandmarketingcouponqrcodegetApiResult 结构体
type AlibabatxcsbrandmarketingcouponqrcodegetApiResult struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 返回内容
	Model *CouponQrcodeResultDo `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
