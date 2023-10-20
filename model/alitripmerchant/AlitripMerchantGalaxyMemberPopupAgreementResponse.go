package alitripmerchant

// AlitripmerchantgalaxymemberpopupagreementResponse 结构体
type AlitripmerchantgalaxymemberpopupagreementResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 用户是否唤起协议弹窗
	Content *UserAgreementVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
