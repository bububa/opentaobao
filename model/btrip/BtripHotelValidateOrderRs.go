package btrip

import (
	"sync"
)

// BtripHotelValidateOrderRs 结构体
type BtripHotelValidateOrderRs struct {
	// 创单Key值，下单时使用
	CreateKey string `json:"create_key,omitempty" xml:"create_key,omitempty"`
	// 优惠信息
	PromotionInfo *BtripHotelPromotionDto `json:"promotion_info,omitempty" xml:"promotion_info,omitempty"`
	// 销售计划id
	RatePlanId int64 `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	// 详细的销售计划
	RatePlanInfo *BtripHotelRatePlanInfoDto `json:"rate_plan_info,omitempty" xml:"rate_plan_info,omitempty"`
}

var poolBtripHotelValidateOrderRs = sync.Pool{
	New: func() any {
		return new(BtripHotelValidateOrderRs)
	},
}

// GetBtripHotelValidateOrderRs() 从对象池中获取BtripHotelValidateOrderRs
func GetBtripHotelValidateOrderRs() *BtripHotelValidateOrderRs {
	return poolBtripHotelValidateOrderRs.Get().(*BtripHotelValidateOrderRs)
}

// ReleaseBtripHotelValidateOrderRs 释放BtripHotelValidateOrderRs
func ReleaseBtripHotelValidateOrderRs(v *BtripHotelValidateOrderRs) {
	v.CreateKey = ""
	v.PromotionInfo = nil
	v.RatePlanId = 0
	v.RatePlanInfo = nil
	poolBtripHotelValidateOrderRs.Put(v)
}
