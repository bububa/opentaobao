package xhotelcrm

import (
	"sync"
)

// SendCouponParam 结构体
type SendCouponParam struct {
	// 会员卡号
	MemberCardNo string `json:"member_card_no,omitempty" xml:"member_card_no,omitempty"`
	// 券有效期截止时间
	CouponEndDate string `json:"coupon_end_date,omitempty" xml:"coupon_end_date,omitempty"`
	// 券实例ID
	CouponInstanceId string `json:"coupon_instance_id,omitempty" xml:"coupon_instance_id,omitempty"`
	// 券有效期开始时间
	CouponStartDate string `json:"coupon_start_date,omitempty" xml:"coupon_start_date,omitempty"`
}

var poolSendCouponParam = sync.Pool{
	New: func() any {
		return new(SendCouponParam)
	},
}

// GetSendCouponParam() 从对象池中获取SendCouponParam
func GetSendCouponParam() *SendCouponParam {
	return poolSendCouponParam.Get().(*SendCouponParam)
}

// ReleaseSendCouponParam 释放SendCouponParam
func ReleaseSendCouponParam(v *SendCouponParam) {
	v.MemberCardNo = ""
	v.CouponEndDate = ""
	v.CouponInstanceId = ""
	v.CouponStartDate = ""
	poolSendCouponParam.Put(v)
}
