package axintrade

import (
	"sync"
)

// HotelOrderValidateRes 结构体
type HotelOrderValidateRes struct {
	// 创建订单需要的key
	CreateKey string `json:"create_key,omitempty" xml:"create_key,omitempty"`
	// 后返佣金费率
	CommissionFeeRate string `json:"commission_fee_rate,omitempty" xml:"commission_fee_rate,omitempty"`
	// 资源标签
	ResourceTag string `json:"resource_tag,omitempty" xml:"resource_tag,omitempty"`
	// 分销模式 FP-底价模式 SP-卖家模式
	DistributeMode string `json:"distribute_mode,omitempty" xml:"distribute_mode,omitempty"`
	// 售卖政策
	RatePlanInfo *RatePlanInfoApiDto `json:"rate_plan_info,omitempty" xml:"rate_plan_info,omitempty"`
	// 售卖政策id
	RatePlanId int64 `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	// 优惠总金额，单位为分
	PromotionTotalPrice int64 `json:"promotion_total_price,omitempty" xml:"promotion_total_price,omitempty"`
	// 1-订单金额酒店开票，分销商开返佣发票 2-订单金额阿信开票，分销商不开票 3-订单金额无票，分销商开返佣发票
	InvoicingMode int64 `json:"invoicing_mode,omitempty" xml:"invoicing_mode,omitempty"`
}

var poolHotelOrderValidateRes = sync.Pool{
	New: func() any {
		return new(HotelOrderValidateRes)
	},
}

// GetHotelOrderValidateRes() 从对象池中获取HotelOrderValidateRes
func GetHotelOrderValidateRes() *HotelOrderValidateRes {
	return poolHotelOrderValidateRes.Get().(*HotelOrderValidateRes)
}

// ReleaseHotelOrderValidateRes 释放HotelOrderValidateRes
func ReleaseHotelOrderValidateRes(v *HotelOrderValidateRes) {
	v.CreateKey = ""
	v.CommissionFeeRate = ""
	v.ResourceTag = ""
	v.DistributeMode = ""
	v.RatePlanInfo = nil
	v.RatePlanId = 0
	v.PromotionTotalPrice = 0
	v.InvoicingMode = 0
	poolHotelOrderValidateRes.Put(v)
}
