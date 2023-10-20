package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeRentTradeBindItemResult 结构体
type AlibabaAlihouseExistinghomeRentTradeBindItemResult struct {
	// 1
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 1
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 1
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 1
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeRentTradeBindItemResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeRentTradeBindItemResult)
	},
}

// GetAlibabaAlihouseExistinghomeRentTradeBindItemResult() 从对象池中获取AlibabaAlihouseExistinghomeRentTradeBindItemResult
func GetAlibabaAlihouseExistinghomeRentTradeBindItemResult() *AlibabaAlihouseExistinghomeRentTradeBindItemResult {
	return poolAlibabaAlihouseExistinghomeRentTradeBindItemResult.Get().(*AlibabaAlihouseExistinghomeRentTradeBindItemResult)
}

// ReleaseAlibabaAlihouseExistinghomeRentTradeBindItemResult 释放AlibabaAlihouseExistinghomeRentTradeBindItemResult
func ReleaseAlibabaAlihouseExistinghomeRentTradeBindItemResult(v *AlibabaAlihouseExistinghomeRentTradeBindItemResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeRentTradeBindItemResult.Put(v)
}
