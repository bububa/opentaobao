package inventory

import (
	"sync"
)

// Strategy 结构体
type Strategy struct {
	// 具体的销售策略，1-先现货库存，后计划库存；2-仅计划库存
	RuleType int64 `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
	// 组合货品情况下，是否支持部分子品现货部分子品计划库存可以进行组合。如果不设置，则都是现货库存，或者都是计划库存才能进行组合
	CombSupport bool `json:"comb_support,omitempty" xml:"comb_support,omitempty"`
}

var poolStrategy = sync.Pool{
	New: func() any {
		return new(Strategy)
	},
}

// GetStrategy() 从对象池中获取Strategy
func GetStrategy() *Strategy {
	return poolStrategy.Get().(*Strategy)
}

// ReleaseStrategy 释放Strategy
func ReleaseStrategy(v *Strategy) {
	v.RuleType = 0
	v.CombSupport = false
	poolStrategy.Put(v)
}
