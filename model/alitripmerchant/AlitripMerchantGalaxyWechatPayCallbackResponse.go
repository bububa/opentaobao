package alitripmerchant

// AlitripMerchantGalaxyWechatPayCallbackResponse 结构体
type AlitripMerchantGalaxyWechatPayCallbackResponse struct {
	// 1
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 1
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 微信要求的返回类
	Content *WechatCallbackResponse `json:"content,omitempty" xml:"content,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
