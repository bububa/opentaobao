package ascp

import (
	"sync"
)

// CapacityRuleUpdateRequest 结构体
type CapacityRuleUpdateRequest struct {
	// 特殊日期产能（组），最多20条
	SpecialDateCapacitys []SpecialDateCapacity `json:"special_date_capacitys,omitempty" xml:"special_date_capacitys>special_date_capacity,omitempty"`
	// 特殊时间段产能（组）
	SpecialTimeCapacitys []SpecialTimeCapacity `json:"special_time_capacitys,omitempty" xml:"special_time_capacitys>special_time_capacity,omitempty"`
	// 业务请求ID（服务商发起请求的ID）
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘天物流服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 服务商仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 通用产能（≥0） 注意：当天可用产能=更新产能-当天已占用产能
	NormalCapacity string `json:"normal_capacity,omitempty" xml:"normal_capacity,omitempty"`
	// 业务请求时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolCapacityRuleUpdateRequest = sync.Pool{
	New: func() any {
		return new(CapacityRuleUpdateRequest)
	},
}

// GetCapacityRuleUpdateRequest() 从对象池中获取CapacityRuleUpdateRequest
func GetCapacityRuleUpdateRequest() *CapacityRuleUpdateRequest {
	return poolCapacityRuleUpdateRequest.Get().(*CapacityRuleUpdateRequest)
}

// ReleaseCapacityRuleUpdateRequest 释放CapacityRuleUpdateRequest
func ReleaseCapacityRuleUpdateRequest(v *CapacityRuleUpdateRequest) {
	v.SpecialDateCapacitys = v.SpecialDateCapacitys[:0]
	v.SpecialTimeCapacitys = v.SpecialTimeCapacitys[:0]
	v.RequestId = ""
	v.SupplierId = ""
	v.WarehouseCode = ""
	v.NormalCapacity = ""
	v.RequestTime = 0
	poolCapacityRuleUpdateRequest.Put(v)
}
