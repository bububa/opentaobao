package tbk

import (
	"sync"
)

// ItemPromotionInfoDto 结构体
type ItemPromotionInfoDto struct {
	// 商品收入比率(%)：商品佣金比率+补贴比率。同物料类接口income_rate
	CommissionRate string `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
	// 佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划，ZX表示自选计划
	CommissionType string `json:"commission_type,omitempty" xml:"commission_type,omitempty"`
	// 商品到手价。单位分
	PromotionPrice int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// 多件价需拍件数
	MultipleItemsPricesCount int64 `json:"multiple_items_prices_count,omitempty" xml:"multiple_items_prices_count,omitempty"`
}

var poolItemPromotionInfoDto = sync.Pool{
	New: func() any {
		return new(ItemPromotionInfoDto)
	},
}

// GetItemPromotionInfoDto() 从对象池中获取ItemPromotionInfoDto
func GetItemPromotionInfoDto() *ItemPromotionInfoDto {
	return poolItemPromotionInfoDto.Get().(*ItemPromotionInfoDto)
}

// ReleaseItemPromotionInfoDto 释放ItemPromotionInfoDto
func ReleaseItemPromotionInfoDto(v *ItemPromotionInfoDto) {
	v.CommissionRate = ""
	v.CommissionType = ""
	v.PromotionPrice = 0
	v.MultipleItemsPricesCount = 0
	poolItemPromotionInfoDto.Put(v)
}
