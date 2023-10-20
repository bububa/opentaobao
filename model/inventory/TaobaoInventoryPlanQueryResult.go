package inventory

import (
	"sync"
)

// TaobaoInventoryPlanQueryResult 结构体
type TaobaoInventoryPlanQueryResult struct {
	// 返回的对象
	Data *PlanInvTopDto `json:"data,omitempty" xml:"data,omitempty"`
	// 返回结果码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

var poolTaobaoInventoryPlanQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoInventoryPlanQueryResult)
	},
}

// GetTaobaoInventoryPlanQueryResult() 从对象池中获取TaobaoInventoryPlanQueryResult
func GetTaobaoInventoryPlanQueryResult() *TaobaoInventoryPlanQueryResult {
	return poolTaobaoInventoryPlanQueryResult.Get().(*TaobaoInventoryPlanQueryResult)
}

// ReleaseTaobaoInventoryPlanQueryResult 释放TaobaoInventoryPlanQueryResult
func ReleaseTaobaoInventoryPlanQueryResult(v *TaobaoInventoryPlanQueryResult) {
	v.Data = nil
	v.ResultCode = nil
	poolTaobaoInventoryPlanQueryResult.Put(v)
}
