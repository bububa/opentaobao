package wdk

import (
	"sync"
)

// AlibabaWdkOrderGetResult 结构体
type AlibabaWdkOrderGetResult struct {
	// 订单列表
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// returnCode
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// returnMsg
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 返回本查询条件下的数据总数
	TotalNumber int64 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

var poolAlibabaWdkOrderGetResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkOrderGetResult)
	},
}

// GetAlibabaWdkOrderGetResult() 从对象池中获取AlibabaWdkOrderGetResult
func GetAlibabaWdkOrderGetResult() *AlibabaWdkOrderGetResult {
	return poolAlibabaWdkOrderGetResult.Get().(*AlibabaWdkOrderGetResult)
}

// ReleaseAlibabaWdkOrderGetResult 释放AlibabaWdkOrderGetResult
func ReleaseAlibabaWdkOrderGetResult(v *AlibabaWdkOrderGetResult) {
	v.Orders = v.Orders[:0]
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.TotalNumber = 0
	poolAlibabaWdkOrderGetResult.Put(v)
}
