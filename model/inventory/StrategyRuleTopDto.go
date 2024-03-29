package inventory

import (
	"sync"
)

// StrategyRuleTopDto 结构体
type StrategyRuleTopDto struct {
	// 1，代表先现货后计划库存；2代表仅卖计划库存
	RuleType int64 `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
	// 是否支持组合情况下，一部分子品用现货，一部分子品用计划库存。true代表支持，空或者false代表  子品都要同类型的库存才能组合起来
	CombSupport bool `json:"comb_support,omitempty" xml:"comb_support,omitempty"`
}

var poolStrategyRuleTopDto = sync.Pool{
	New: func() any {
		return new(StrategyRuleTopDto)
	},
}

// GetStrategyRuleTopDto() 从对象池中获取StrategyRuleTopDto
func GetStrategyRuleTopDto() *StrategyRuleTopDto {
	return poolStrategyRuleTopDto.Get().(*StrategyRuleTopDto)
}

// ReleaseStrategyRuleTopDto 释放StrategyRuleTopDto
func ReleaseStrategyRuleTopDto(v *StrategyRuleTopDto) {
	v.RuleType = 0
	v.CombSupport = false
	poolStrategyRuleTopDto.Put(v)
}
