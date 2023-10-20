package trade

import (
	"sync"
)

// OrderResult 结构体
type OrderResult struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 取消的订单
	Result *OrderObject `json:"result,omitempty" xml:"result,omitempty"`
	// 创建的订单
	Trade *TradeOrder `json:"trade,omitempty" xml:"trade,omitempty"`
	// 是否取消成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolOrderResult = sync.Pool{
	New: func() any {
		return new(OrderResult)
	},
}

// GetOrderResult() 从对象池中获取OrderResult
func GetOrderResult() *OrderResult {
	return poolOrderResult.Get().(*OrderResult)
}

// ReleaseOrderResult 释放OrderResult
func ReleaseOrderResult(v *OrderResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Result = nil
	v.Trade = nil
	v.Success = false
	poolOrderResult.Put(v)
}
