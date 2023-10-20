package promotion

import (
	"sync"
)

// CouponTemplateOperateRequest 结构体
type CouponTemplateOperateRequest struct {
	// 券模版
	CouponTemplate *CouponTemplate `json:"coupon_template,omitempty" xml:"coupon_template,omitempty"`
	// 用户信息
	UserInfo *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
}

var poolCouponTemplateOperateRequest = sync.Pool{
	New: func() any {
		return new(CouponTemplateOperateRequest)
	},
}

// GetCouponTemplateOperateRequest() 从对象池中获取CouponTemplateOperateRequest
func GetCouponTemplateOperateRequest() *CouponTemplateOperateRequest {
	return poolCouponTemplateOperateRequest.Get().(*CouponTemplateOperateRequest)
}

// ReleaseCouponTemplateOperateRequest 释放CouponTemplateOperateRequest
func ReleaseCouponTemplateOperateRequest(v *CouponTemplateOperateRequest) {
	v.CouponTemplate = nil
	v.UserInfo = nil
	poolCouponTemplateOperateRequest.Put(v)
}
