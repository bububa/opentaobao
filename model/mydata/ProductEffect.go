package mydata

import (
	"sync"
)

// ProductEffect 结构体
type ProductEffect struct {
	// 产品效果
	Effects []EffectEntity `json:"effects,omitempty" xml:"effects>effect_entity,omitempty"`
	// 返回结果数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolProductEffect = sync.Pool{
	New: func() any {
		return new(ProductEffect)
	},
}

// GetProductEffect() 从对象池中获取ProductEffect
func GetProductEffect() *ProductEffect {
	return poolProductEffect.Get().(*ProductEffect)
}

// ReleaseProductEffect 释放ProductEffect
func ReleaseProductEffect(v *ProductEffect) {
	v.Effects = v.Effects[:0]
	v.TotalCount = 0
	poolProductEffect.Put(v)
}
