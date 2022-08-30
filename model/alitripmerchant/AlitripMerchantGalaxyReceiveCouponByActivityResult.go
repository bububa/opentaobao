package alitripmerchant

// AlitripMerchantGalaxyReceiveCouponByActivityResult 结构体
type AlitripMerchantGalaxyReceiveCouponByActivityResult struct {
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否参与活动成功
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}
