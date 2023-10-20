package alitripmerchant

// AlitripmerchantgalaxycouponvalidlistResponse 结构体
type AlitripmerchantgalaxycouponvalidlistResponse struct {
	// 状态码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 优惠券实例
	Content *WeChatCouponVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
