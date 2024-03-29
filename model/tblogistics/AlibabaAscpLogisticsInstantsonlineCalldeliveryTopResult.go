package tblogistics

import (
	"sync"
)

// AlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult 结构体
type AlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult struct {
	// 取号流水号
	CwOrderId string `json:"cw_order_id,omitempty" xml:"cw_order_id,omitempty"`
	// 物流单号
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
	// 资源CODE
	ResourceCode string `json:"resource_code,omitempty" xml:"resource_code,omitempty"`
	// 取件码
	PickupCode string `json:"pickup_code,omitempty" xml:"pickup_code,omitempty"`
	// 扩展
	Features string `json:"features,omitempty" xml:"features,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult)
	},
}

// GetAlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult() 从对象池中获取AlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult
func GetAlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult() *AlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult {
	return poolAlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult.Get().(*AlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult)
}

// ReleaseAlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult 释放AlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult
func ReleaseAlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult(v *AlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult) {
	v.CwOrderId = ""
	v.WaybillCode = ""
	v.ResourceCode = ""
	v.PickupCode = ""
	v.Features = ""
	v.Success = false
	poolAlibabaAscpLogisticsInstantsonlineCalldeliveryTopResult.Put(v)
}
