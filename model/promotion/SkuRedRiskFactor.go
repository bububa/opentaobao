package promotion

import (
	"sync"
)

// SkuRedRiskFactor 结构体
type SkuRedRiskFactor struct {
	// skuId
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// sku红线价格
	AmountAt int64 `json:"amount_at,omitempty" xml:"amount_at,omitempty"`
}

var poolSkuRedRiskFactor = sync.Pool{
	New: func() any {
		return new(SkuRedRiskFactor)
	},
}

// GetSkuRedRiskFactor() 从对象池中获取SkuRedRiskFactor
func GetSkuRedRiskFactor() *SkuRedRiskFactor {
	return poolSkuRedRiskFactor.Get().(*SkuRedRiskFactor)
}

// ReleaseSkuRedRiskFactor 释放SkuRedRiskFactor
func ReleaseSkuRedRiskFactor(v *SkuRedRiskFactor) {
	v.SkuId = ""
	v.AmountAt = 0
	poolSkuRedRiskFactor.Put(v)
}
