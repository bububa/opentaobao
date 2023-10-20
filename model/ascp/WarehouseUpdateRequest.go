package ascp

import (
	"sync"
)

// WarehouseUpdateRequest 结构体
type WarehouseUpdateRequest struct {
	// 仓信息数组
	WarehouseInfos []WarehouseInfo `json:"warehouse_infos,omitempty" xml:"warehouse_infos>warehouse_info,omitempty"`
	// 业务请求ID（服务商发起请求的ID）
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘天物流服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 业务请求时间戳（毫秒）
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolWarehouseUpdateRequest = sync.Pool{
	New: func() any {
		return new(WarehouseUpdateRequest)
	},
}

// GetWarehouseUpdateRequest() 从对象池中获取WarehouseUpdateRequest
func GetWarehouseUpdateRequest() *WarehouseUpdateRequest {
	return poolWarehouseUpdateRequest.Get().(*WarehouseUpdateRequest)
}

// ReleaseWarehouseUpdateRequest 释放WarehouseUpdateRequest
func ReleaseWarehouseUpdateRequest(v *WarehouseUpdateRequest) {
	v.WarehouseInfos = v.WarehouseInfos[:0]
	v.RequestId = ""
	v.SupplierId = ""
	v.RequestTime = 0
	poolWarehouseUpdateRequest.Put(v)
}
