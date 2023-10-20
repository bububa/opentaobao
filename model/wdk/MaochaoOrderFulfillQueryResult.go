package wdk

import (
	"sync"
)

// MaochaoOrderFulfillQueryResult 结构体
type MaochaoOrderFulfillQueryResult struct {
	// 履约单列表
	FulfillOrderList []MaochaoWdkOrderFulfillDto `json:"fulfill_order_list,omitempty" xml:"fulfill_order_list>maochao_wdk_order_fulfill_dto,omitempty"`
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回码说明
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolMaochaoOrderFulfillQueryResult = sync.Pool{
	New: func() any {
		return new(MaochaoOrderFulfillQueryResult)
	},
}

// GetMaochaoOrderFulfillQueryResult() 从对象池中获取MaochaoOrderFulfillQueryResult
func GetMaochaoOrderFulfillQueryResult() *MaochaoOrderFulfillQueryResult {
	return poolMaochaoOrderFulfillQueryResult.Get().(*MaochaoOrderFulfillQueryResult)
}

// ReleaseMaochaoOrderFulfillQueryResult 释放MaochaoOrderFulfillQueryResult
func ReleaseMaochaoOrderFulfillQueryResult(v *MaochaoOrderFulfillQueryResult) {
	v.FulfillOrderList = v.FulfillOrderList[:0]
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.Success = false
	poolMaochaoOrderFulfillQueryResult.Put(v)
}
