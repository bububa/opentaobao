package alitripmerchant

// WeChatCouponVO 结构体
type WeChatCouponVO struct {
	// 优惠券实例对象
	CouponInfoVO []CouponInfoVO `json:"coupon_info_v_o,omitempty" xml:"coupon_info_v_o>coupon_info_vo,omitempty"`
	// 1
	SendCouponParams []SendCouponParams `json:"send_coupon_params,omitempty" xml:"send_coupon_params>send_coupon_params,omitempty"`
	// 加密串
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 商户号
	SendCouponMerchant string `json:"send_coupon_merchant,omitempty" xml:"send_coupon_merchant,omitempty"`
}
