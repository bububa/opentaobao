package xhotel

import (
	"sync"
)

// UpdatePromoParam 结构体
type UpdatePromoParam struct {
	// 营销活动关联的价格计划和房型
	RateInfos []RateInfo `json:"rate_infos,omitempty" xml:"rate_infos>rate_info,omitempty"`
	// 外部营销活动的code，最长40个字符
	OuterActivityCode string `json:"outer_activity_code,omitempty" xml:"outer_activity_code,omitempty"`
	// 营销活动的具体参数对象，在每次添加更新的时候，long_order_info、early_booking_info、daily_booking_info 只能填1种类型，其他2种类型为空
	PromoInfo *PromoInfo `json:"promo_info,omitempty" xml:"promo_info,omitempty"`
}

var poolUpdatePromoParam = sync.Pool{
	New: func() any {
		return new(UpdatePromoParam)
	},
}

// GetUpdatePromoParam() 从对象池中获取UpdatePromoParam
func GetUpdatePromoParam() *UpdatePromoParam {
	return poolUpdatePromoParam.Get().(*UpdatePromoParam)
}

// ReleaseUpdatePromoParam 释放UpdatePromoParam
func ReleaseUpdatePromoParam(v *UpdatePromoParam) {
	v.RateInfos = v.RateInfos[:0]
	v.OuterActivityCode = ""
	v.PromoInfo = nil
	poolUpdatePromoParam.Put(v)
}
