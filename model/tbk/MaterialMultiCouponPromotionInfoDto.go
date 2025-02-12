package tbk

import (
	"sync"
)

// MaterialMultiCouponPromotionInfoDto 结构体
type MaterialMultiCouponPromotionInfoDto struct {
	// 优惠名称
	CouponTitle string `json:"coupon_title,omitempty" xml:"coupon_title,omitempty"`
	// 优惠券信息，满XX元减XX元，满x件减y元
	CouponDesc string `json:"coupon_desc,omitempty" xml:"coupon_desc,omitempty"`
	// 优惠金额（元）
	CoupoonFee string `json:"coupoon_fee,omitempty" xml:"coupoon_fee,omitempty"`
	// 优惠开始时间
	CouponStartTime string `json:"coupon_start_time,omitempty" xml:"coupon_start_time,omitempty"`
	// 优惠结束时间
	CouponEndTime string `json:"coupon_end_time,omitempty" xml:"coupon_end_time,omitempty"`
	// 优惠ID
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 优惠剩余数量
	CouponRemainCount int64 `json:"coupon_remain_count,omitempty" xml:"coupon_remain_count,omitempty"`
}

var poolMaterialMultiCouponPromotionInfoDto = sync.Pool{
	New: func() any {
		return new(MaterialMultiCouponPromotionInfoDto)
	},
}

// GetMaterialMultiCouponPromotionInfoDto() 从对象池中获取MaterialMultiCouponPromotionInfoDto
func GetMaterialMultiCouponPromotionInfoDto() *MaterialMultiCouponPromotionInfoDto {
	return poolMaterialMultiCouponPromotionInfoDto.Get().(*MaterialMultiCouponPromotionInfoDto)
}

// ReleaseMaterialMultiCouponPromotionInfoDto 释放MaterialMultiCouponPromotionInfoDto
func ReleaseMaterialMultiCouponPromotionInfoDto(v *MaterialMultiCouponPromotionInfoDto) {
	v.CouponTitle = ""
	v.CouponDesc = ""
	v.CoupoonFee = ""
	v.CouponStartTime = ""
	v.CouponEndTime = ""
	v.ActivityId = ""
	v.CouponRemainCount = 0
	poolMaterialMultiCouponPromotionInfoDto.Put(v)
}
