package ascp

import (
	"sync"
)

// WarehouseOperationUpdateRequest 结构体
type WarehouseOperationUpdateRequest struct {
	// 业务请求ID（服务商发起请求的ID）
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘天物流服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 服务商仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 业务请求时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 1-通用规则；2-个性化规则
	RuleType int64 `json:"rule_type,omitempty" xml:"rule_type,omitempty"`
	// 仓时效规则信息，通用规则传入时必填
	NormalWarehouseTimingRule *NormalWarehouseTimingRule `json:"normal_warehouse_timing_rule,omitempty" xml:"normal_warehouse_timing_rule,omitempty"`
	// 个性化效规则信息（组），个性化规则传入时必填，最多20个货主
	SpecialWarehouseTimingRules *SpecialWarehouseTimingRules `json:"special_warehouse_timing_rules,omitempty" xml:"special_warehouse_timing_rules,omitempty"`
}

var poolWarehouseOperationUpdateRequest = sync.Pool{
	New: func() any {
		return new(WarehouseOperationUpdateRequest)
	},
}

// GetWarehouseOperationUpdateRequest() 从对象池中获取WarehouseOperationUpdateRequest
func GetWarehouseOperationUpdateRequest() *WarehouseOperationUpdateRequest {
	return poolWarehouseOperationUpdateRequest.Get().(*WarehouseOperationUpdateRequest)
}

// ReleaseWarehouseOperationUpdateRequest 释放WarehouseOperationUpdateRequest
func ReleaseWarehouseOperationUpdateRequest(v *WarehouseOperationUpdateRequest) {
	v.RequestId = ""
	v.SupplierId = ""
	v.WarehouseCode = ""
	v.RequestTime = 0
	v.RuleType = 0
	v.NormalWarehouseTimingRule = nil
	v.SpecialWarehouseTimingRules = nil
	poolWarehouseOperationUpdateRequest.Put(v)
}
