package cainiaohandover

import (
	"sync"
)

// OpenParcelOrderQueryResponse 结构体
type OpenParcelOrderQueryResponse struct {
	// 交接仓编码，快递揽收场景,大包交接目的地国际分拨
	HandoverWarehouseCode string `json:"handover_warehouse_code,omitempty" xml:"handover_warehouse_code,omitempty"`
	// 交接仓名称，快递揽收场景,大包交接目的地国际分拨
	HandoverWarehouseName string `json:"handover_warehouse_name,omitempty" xml:"handover_warehouse_name,omitempty"`
	// 关联的大包的编码
	HandoverContentCode string `json:"handover_content_code,omitempty" xml:"handover_content_code,omitempty"`
	// 关联的交接单ID
	HandoverOrderId int64 `json:"handover_order_id,omitempty" xml:"handover_order_id,omitempty"`
	// 关联的大包ID
	HandoverContentId int64 `json:"handover_content_id,omitempty" xml:"handover_content_id,omitempty"`
	// 该小包是否已经组包
	HasBeenHandover bool `json:"has_been_handover,omitempty" xml:"has_been_handover,omitempty"`
	// 是否能组包
	CanCreateHandover bool `json:"can_create_handover,omitempty" xml:"can_create_handover,omitempty"`
}

var poolOpenParcelOrderQueryResponse = sync.Pool{
	New: func() any {
		return new(OpenParcelOrderQueryResponse)
	},
}

// GetOpenParcelOrderQueryResponse() 从对象池中获取OpenParcelOrderQueryResponse
func GetOpenParcelOrderQueryResponse() *OpenParcelOrderQueryResponse {
	return poolOpenParcelOrderQueryResponse.Get().(*OpenParcelOrderQueryResponse)
}

// ReleaseOpenParcelOrderQueryResponse 释放OpenParcelOrderQueryResponse
func ReleaseOpenParcelOrderQueryResponse(v *OpenParcelOrderQueryResponse) {
	v.HandoverWarehouseCode = ""
	v.HandoverWarehouseName = ""
	v.HandoverContentCode = ""
	v.HandoverOrderId = 0
	v.HandoverContentId = 0
	v.HasBeenHandover = false
	v.CanCreateHandover = false
	poolOpenParcelOrderQueryResponse.Put(v)
}
