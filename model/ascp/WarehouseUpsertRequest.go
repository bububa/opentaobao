package ascp

import (
	"sync"
)

// WarehouseUpsertRequest 结构体
type WarehouseUpsertRequest struct {
	// 仓信息数组 ；最多50条
	WarehouseInfos []WarehouseInfo `json:"warehouse_infos,omitempty" xml:"warehouse_infos>warehouse_info,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolWarehouseUpsertRequest = sync.Pool{
	New: func() any {
		return new(WarehouseUpsertRequest)
	},
}

// GetWarehouseUpsertRequest() 从对象池中获取WarehouseUpsertRequest
func GetWarehouseUpsertRequest() *WarehouseUpsertRequest {
	return poolWarehouseUpsertRequest.Get().(*WarehouseUpsertRequest)
}

// ReleaseWarehouseUpsertRequest 释放WarehouseUpsertRequest
func ReleaseWarehouseUpsertRequest(v *WarehouseUpsertRequest) {
	v.WarehouseInfos = v.WarehouseInfos[:0]
	v.RequestId = ""
	v.RequestTime = 0
	poolWarehouseUpsertRequest.Put(v)
}
