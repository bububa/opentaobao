package wdk

import (
	"sync"
)

// OrderQueryResult 结构体
type OrderQueryResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 订单数据
	Result *OrderSuccessResponse `json:"result,omitempty" xml:"result,omitempty"`
	// 接口状态
	State bool `json:"state,omitempty" xml:"state,omitempty"`
}

var poolOrderQueryResult = sync.Pool{
	New: func() any {
		return new(OrderQueryResult)
	},
}

// GetOrderQueryResult() 从对象池中获取OrderQueryResult
func GetOrderQueryResult() *OrderQueryResult {
	return poolOrderQueryResult.Get().(*OrderQueryResult)
}

// ReleaseOrderQueryResult 释放OrderQueryResult
func ReleaseOrderQueryResult(v *OrderQueryResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Result = nil
	v.State = false
	poolOrderQueryResult.Put(v)
}
