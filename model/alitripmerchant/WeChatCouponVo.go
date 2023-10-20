package alitripmerchant

import (
	"sync"
)

// WeChatCouponVo 结构体
type WeChatCouponVo struct {
	// 优惠券实例对象
	CouponInfoVO []CouponInfoVo `json:"coupon_info_v_o,omitempty" xml:"coupon_info_v_o>coupon_info_vo,omitempty"`
	// 1
	SendCouponParams []SendCouponParams `json:"send_coupon_params,omitempty" xml:"send_coupon_params>send_coupon_params,omitempty"`
	// 加密串
	Sign string `json:"sign,omitempty" xml:"sign,omitempty"`
	// 商户号
	SendCouponMerchant string `json:"send_coupon_merchant,omitempty" xml:"send_coupon_merchant,omitempty"`
}

var poolWeChatCouponVo = sync.Pool{
	New: func() any {
		return new(WeChatCouponVo)
	},
}

// GetWeChatCouponVo() 从对象池中获取WeChatCouponVo
func GetWeChatCouponVo() *WeChatCouponVo {
	return poolWeChatCouponVo.Get().(*WeChatCouponVo)
}

// ReleaseWeChatCouponVo 释放WeChatCouponVo
func ReleaseWeChatCouponVo(v *WeChatCouponVo) {
	v.CouponInfoVO = v.CouponInfoVO[:0]
	v.SendCouponParams = v.SendCouponParams[:0]
	v.Sign = ""
	v.SendCouponMerchant = ""
	poolWeChatCouponVo.Put(v)
}
