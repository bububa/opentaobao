package tblogistics

import (
	"sync"
)

// AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult 结构体
type AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult struct {
	// 数据
	Data *PriorCallDeliveryTopResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult)
	},
}

// GetAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult() 从对象池中获取AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult
func GetAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult() *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult {
	return poolAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult.Get().(*AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult)
}

// ReleaseAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult 释放AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult
func ReleaseAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult(v *AlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult) {
	v.Data = nil
	v.Success = false
	poolAlibabaAscpLogisticsInstantsonlinePriorcalldeliveryTopResult.Put(v)
}
