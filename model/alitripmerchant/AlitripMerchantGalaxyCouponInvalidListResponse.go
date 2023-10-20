package alitripmerchant

// AlitripmerchantgalaxycouponinvalidlistResponse 结构体
type AlitripmerchantgalaxycouponinvalidlistResponse struct {
	// 优惠券详情
	Contents []CouponInfo `json:"contents,omitempty" xml:"contents>coupon_info,omitempty"`
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
