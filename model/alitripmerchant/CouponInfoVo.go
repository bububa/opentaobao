package alitripmerchant

import (
	"sync"
)

// CouponInfoVo 结构体
type CouponInfoVo struct {
	// 折扣
	DiscountAmount string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 门槛
	ConditionAmount string `json:"condition_amount,omitempty" xml:"condition_amount,omitempty"`
	// 描述
	DetailDesc string `json:"detail_desc,omitempty" xml:"detail_desc,omitempty"`
	// 优惠券名字
	CouponName string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
	// 过期时间
	ExpiredTimeMin string `json:"expired_time_min,omitempty" xml:"expired_time_min,omitempty"`
	// 优惠券类型
	CouponType string `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
	// 描述
	ConditionDesc string `json:"condition_desc,omitempty" xml:"condition_desc,omitempty"`
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 实例id
	InstanceId int64 `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// 模板id
	CouponTemplateId int64 `json:"coupon_template_id,omitempty" xml:"coupon_template_id,omitempty"`
}

var poolCouponInfoVo = sync.Pool{
	New: func() any {
		return new(CouponInfoVo)
	},
}

// GetCouponInfoVo() 从对象池中获取CouponInfoVo
func GetCouponInfoVo() *CouponInfoVo {
	return poolCouponInfoVo.Get().(*CouponInfoVo)
}

// ReleaseCouponInfoVo 释放CouponInfoVo
func ReleaseCouponInfoVo(v *CouponInfoVo) {
	v.DiscountAmount = ""
	v.ConditionAmount = ""
	v.DetailDesc = ""
	v.CouponName = ""
	v.ExpiredTimeMin = ""
	v.CouponType = ""
	v.ConditionDesc = ""
	v.Status = ""
	v.InstanceId = 0
	v.CouponTemplateId = 0
	poolCouponInfoVo.Put(v)
}
