package tbk

import (
	"sync"
)

// PromotionInfoDto 结构体
type PromotionInfoDto struct {
	// 商品收入比率(%)：商品佣金比率+补贴比率。同物料类接口income_rate
	CommissionRate string `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
	// 佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划，ZX表示自选计划
	CommissionType string `json:"commission_type,omitempty" xml:"commission_type,omitempty"`
	// 多件价需拍件数
	MultipleItemsPricesCount int64 `json:"multiple_items_prices_count,omitempty" xml:"multiple_items_prices_count,omitempty"`
	// 商品到手价。单位分
	PromotionPrice int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
}

var poolPromotionInfoDto = sync.Pool{
	New: func() any {
		return new(PromotionInfoDto)
	},
}

// GetPromotionInfoDto() 从对象池中获取PromotionInfoDto
func GetPromotionInfoDto() *PromotionInfoDto {
	return poolPromotionInfoDto.Get().(*PromotionInfoDto)
}

// ReleasePromotionInfoDto 释放PromotionInfoDto
func ReleasePromotionInfoDto(v *PromotionInfoDto) {
	v.CommissionRate = ""
	v.CommissionType = ""
	v.MultipleItemsPricesCount = 0
	v.PromotionPrice = 0
	poolPromotionInfoDto.Put(v)
}
