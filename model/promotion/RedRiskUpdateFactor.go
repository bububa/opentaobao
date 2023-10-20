package promotion

import (
	"sync"
)

// RedRiskUpdateFactor 结构体
type RedRiskUpdateFactor struct {
	// 需要删除的sku红线价格
	RemoveSkuIds []int64 `json:"remove_sku_ids,omitempty" xml:"remove_sku_ids>int64,omitempty"`
	// 新增sku级别红线价格
	SkuRiskFactors []SkuRedRiskFactor `json:"sku_risk_factors,omitempty" xml:"sku_risk_factors>sku_red_risk_factor,omitempty"`
	// 风险等级设置
	RiskLevels []RiskLevelParam `json:"risk_levels,omitempty" xml:"risk_levels>risk_level_param,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品红线价格
	AmountAt int64 `json:"amount_at,omitempty" xml:"amount_at,omitempty"`
}

var poolRedRiskUpdateFactor = sync.Pool{
	New: func() any {
		return new(RedRiskUpdateFactor)
	},
}

// GetRedRiskUpdateFactor() 从对象池中获取RedRiskUpdateFactor
func GetRedRiskUpdateFactor() *RedRiskUpdateFactor {
	return poolRedRiskUpdateFactor.Get().(*RedRiskUpdateFactor)
}

// ReleaseRedRiskUpdateFactor 释放RedRiskUpdateFactor
func ReleaseRedRiskUpdateFactor(v *RedRiskUpdateFactor) {
	v.RemoveSkuIds = v.RemoveSkuIds[:0]
	v.SkuRiskFactors = v.SkuRiskFactors[:0]
	v.RiskLevels = v.RiskLevels[:0]
	v.ItemId = 0
	v.AmountAt = 0
	poolRedRiskUpdateFactor.Put(v)
}
