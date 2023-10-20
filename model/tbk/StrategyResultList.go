package tbk

import (
	"sync"
)

// StrategyResultList 结构体
type StrategyResultList struct {
	// 策略ID
	StrategyId string `json:"strategy_id,omitempty" xml:"strategy_id,omitempty"`
	// 状态：1-符合活动要求 ，3-用户不匹配活动，4-系统异常，6-策略ID不存在，7-策略ID无效
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolStrategyResultList = sync.Pool{
	New: func() any {
		return new(StrategyResultList)
	},
}

// GetStrategyResultList() 从对象池中获取StrategyResultList
func GetStrategyResultList() *StrategyResultList {
	return poolStrategyResultList.Get().(*StrategyResultList)
}

// ReleaseStrategyResultList 释放StrategyResultList
func ReleaseStrategyResultList(v *StrategyResultList) {
	v.StrategyId = ""
	v.Status = ""
	poolStrategyResultList.Put(v)
}
