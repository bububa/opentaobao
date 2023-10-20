package alitripmerchant

// AlitripmerchantgalaxyactivitycouponlistResponse 结构体
type AlitripmerchantgalaxyactivitycouponlistResponse struct {
	// 活动详情
	Contents []CouponActivity `json:"contents,omitempty" xml:"contents>coupon_activity,omitempty"`
	// 错误信息
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
