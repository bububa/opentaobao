package idle

import (
	"sync"
)

// AlibabaIdleRentOrderLogisticsDeliverTopResult 结构体
type AlibabaIdleRentOrderLogisticsDeliverTopResult struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// data
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaIdleRentOrderLogisticsDeliverTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaIdleRentOrderLogisticsDeliverTopResult)
	},
}

// GetAlibabaIdleRentOrderLogisticsDeliverTopResult() 从对象池中获取AlibabaIdleRentOrderLogisticsDeliverTopResult
func GetAlibabaIdleRentOrderLogisticsDeliverTopResult() *AlibabaIdleRentOrderLogisticsDeliverTopResult {
	return poolAlibabaIdleRentOrderLogisticsDeliverTopResult.Get().(*AlibabaIdleRentOrderLogisticsDeliverTopResult)
}

// ReleaseAlibabaIdleRentOrderLogisticsDeliverTopResult 释放AlibabaIdleRentOrderLogisticsDeliverTopResult
func ReleaseAlibabaIdleRentOrderLogisticsDeliverTopResult(v *AlibabaIdleRentOrderLogisticsDeliverTopResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = false
	v.Success = false
	poolAlibabaIdleRentOrderLogisticsDeliverTopResult.Put(v)
}
