package tblogistics

import (
	"sync"
)

// AlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult 结构体
type AlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult struct {
	// 数据
	Data *GetDeliveryOrderTopResponse `json:"data,omitempty" xml:"data,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult)
	},
}

// GetAlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult() 从对象池中获取AlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult
func GetAlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult() *AlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult {
	return poolAlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult.Get().(*AlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult)
}

// ReleaseAlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult 释放AlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult
func ReleaseAlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult(v *AlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult) {
	v.Data = nil
	v.Success = false
	poolAlibabaAscpLogisticsInstantsonlineDeliveryorderGetTopResult.Put(v)
}
