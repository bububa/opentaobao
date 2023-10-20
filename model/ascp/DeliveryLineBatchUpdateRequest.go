package ascp

import (
	"sync"
)

// DeliveryLineBatchUpdateRequest 结构体
type DeliveryLineBatchUpdateRequest struct {
	// 通用到货线路规则（组）
	NormalSignLineRules []NormalSignLineRule `json:"normal_sign_line_rules,omitempty" xml:"normal_sign_line_rules>normal_sign_line_rule,omitempty"`
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
	// 个性化到货线路规则（组）；
	SpecialSignLineRules *SpecialSignLineRules `json:"special_sign_line_rules,omitempty" xml:"special_sign_line_rules,omitempty"`
}

var poolDeliveryLineBatchUpdateRequest = sync.Pool{
	New: func() any {
		return new(DeliveryLineBatchUpdateRequest)
	},
}

// GetDeliveryLineBatchUpdateRequest() 从对象池中获取DeliveryLineBatchUpdateRequest
func GetDeliveryLineBatchUpdateRequest() *DeliveryLineBatchUpdateRequest {
	return poolDeliveryLineBatchUpdateRequest.Get().(*DeliveryLineBatchUpdateRequest)
}

// ReleaseDeliveryLineBatchUpdateRequest 释放DeliveryLineBatchUpdateRequest
func ReleaseDeliveryLineBatchUpdateRequest(v *DeliveryLineBatchUpdateRequest) {
	v.NormalSignLineRules = v.NormalSignLineRules[:0]
	v.RequestId = ""
	v.SupplierId = ""
	v.WarehouseCode = ""
	v.RequestTime = 0
	v.RuleType = 0
	v.SpecialSignLineRules = nil
	poolDeliveryLineBatchUpdateRequest.Put(v)
}
