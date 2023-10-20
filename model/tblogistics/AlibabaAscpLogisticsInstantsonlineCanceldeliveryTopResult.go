package tblogistics

import (
	"sync"
)

// AlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult 结构体
type AlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult struct {
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否已取消
	Canceled bool `json:"canceled,omitempty" xml:"canceled,omitempty"`
}

var poolAlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult)
	},
}

// GetAlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult() 从对象池中获取AlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult
func GetAlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult() *AlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult {
	return poolAlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult.Get().(*AlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult)
}

// ReleaseAlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult 释放AlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult
func ReleaseAlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult(v *AlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult) {
	v.Success = false
	v.Canceled = false
	poolAlibabaAscpLogisticsInstantsonlineCanceldeliveryTopResult.Put(v)
}
