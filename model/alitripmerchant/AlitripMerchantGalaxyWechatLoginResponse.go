package alitripmerchant

// AlitripMerchantGalaxyWechatLoginResponse 结构体
type AlitripMerchantGalaxyWechatLoginResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
