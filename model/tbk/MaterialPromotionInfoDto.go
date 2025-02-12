package tbk

import (
	"sync"
)

// MaterialPromotionInfoDto 结构体
type MaterialPromotionInfoDto struct {
	// 商品收入比率(%)：商品佣金比率+补贴比率。同物料类接口income_rate
	CommissionRate string `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
	// 佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划，ZX表示自选计划
	CommissionType string `json:"commission_type,omitempty" xml:"commission_type,omitempty"`
	// 多件价需拍件数
	MultipleItemsPricesCount int64 `json:"multiple_items_prices_count,omitempty" xml:"multiple_items_prices_count,omitempty"`
}

var poolMaterialPromotionInfoDto = sync.Pool{
	New: func() any {
		return new(MaterialPromotionInfoDto)
	},
}

// GetMaterialPromotionInfoDto() 从对象池中获取MaterialPromotionInfoDto
func GetMaterialPromotionInfoDto() *MaterialPromotionInfoDto {
	return poolMaterialPromotionInfoDto.Get().(*MaterialPromotionInfoDto)
}

// ReleaseMaterialPromotionInfoDto 释放MaterialPromotionInfoDto
func ReleaseMaterialPromotionInfoDto(v *MaterialPromotionInfoDto) {
	v.CommissionRate = ""
	v.CommissionType = ""
	v.MultipleItemsPricesCount = 0
	poolMaterialPromotionInfoDto.Put(v)
}
