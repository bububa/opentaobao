package alitripmerchant

// AlitripMerchantGalaxyCouponValidListResponse 结构体
type AlitripMerchantGalaxyCouponValidListResponse struct {
	// 状态码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 优惠券实例
	Content *WeChatCouponVO `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
