package tbk

import (
	"sync"
)

// CouponInfoDto 结构体
type CouponInfoDto struct {
	// 优惠券结束时间
	CouponEndTime string `json:"coupon_end_time,omitempty" xml:"coupon_end_time,omitempty"`
	// 券ID
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 券面额
	CouponAmount string `json:"coupon_amount,omitempty" xml:"coupon_amount,omitempty"`
	// 优惠券开始时间
	CouponStartTime string `json:"coupon_start_time,omitempty" xml:"coupon_start_time,omitempty"`
	// 优惠券信息，满XX元减XX元
	CouponDesc string `json:"coupon_desc,omitempty" xml:"coupon_desc,omitempty"`
	// 优惠券剩余数量
	CouponRemainCount int64 `json:"coupon_remain_count,omitempty" xml:"coupon_remain_count,omitempty"`
	//  0-店铺 1-单品
	CouponType int64 `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
}

var poolCouponInfoDto = sync.Pool{
	New: func() any {
		return new(CouponInfoDto)
	},
}

// GetCouponInfoDto() 从对象池中获取CouponInfoDto
func GetCouponInfoDto() *CouponInfoDto {
	return poolCouponInfoDto.Get().(*CouponInfoDto)
}

// ReleaseCouponInfoDto 释放CouponInfoDto
func ReleaseCouponInfoDto(v *CouponInfoDto) {
	v.CouponEndTime = ""
	v.ActivityId = ""
	v.CouponAmount = ""
	v.CouponStartTime = ""
	v.CouponDesc = ""
	v.CouponRemainCount = 0
	v.CouponType = 0
	poolCouponInfoDto.Put(v)
}
