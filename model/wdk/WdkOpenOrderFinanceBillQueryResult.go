package wdk

import (
	"sync"
)

// WdkOpenOrderFinanceBillQueryResult 结构体
type WdkOpenOrderFinanceBillQueryResult struct {
	// 账单列表
	Bills []WdkOpenOrderFinanceBillDo `json:"bills,omitempty" xml:"bills>wdk_open_order_finance_bill_do,omitempty"`
	// 结果信息
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 结果码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 总数量，只在查询第一页时返回
	TotalNumber int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
	// 下一页查询的入参，当为-1时表示没有更多数据
	NextId int64 `json:"next_id,omitempty" xml:"next_id,omitempty"`
	// 成功或失败，调用方需要根据该状态判断是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolWdkOpenOrderFinanceBillQueryResult = sync.Pool{
	New: func() any {
		return new(WdkOpenOrderFinanceBillQueryResult)
	},
}

// GetWdkOpenOrderFinanceBillQueryResult() 从对象池中获取WdkOpenOrderFinanceBillQueryResult
func GetWdkOpenOrderFinanceBillQueryResult() *WdkOpenOrderFinanceBillQueryResult {
	return poolWdkOpenOrderFinanceBillQueryResult.Get().(*WdkOpenOrderFinanceBillQueryResult)
}

// ReleaseWdkOpenOrderFinanceBillQueryResult 释放WdkOpenOrderFinanceBillQueryResult
func ReleaseWdkOpenOrderFinanceBillQueryResult(v *WdkOpenOrderFinanceBillQueryResult) {
	v.Bills = v.Bills[:0]
	v.ReturnMsg = ""
	v.ReturnCode = ""
	v.TotalNumber = 0
	v.NextId = 0
	v.Success = false
	poolWdkOpenOrderFinanceBillQueryResult.Put(v)
}
