package tmallservice

import (
	"sync"
)

// ReassignStoreRequest 结构体
type ReassignStoreRequest struct {
	// 工单id
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
	// 目标门店id
	TargetServiceStoreId int64 `json:"target_service_store_id,omitempty" xml:"target_service_store_id,omitempty"`
	// 原门店id
	OriginServiceStoreId int64 `json:"origin_service_store_id,omitempty" xml:"origin_service_store_id,omitempty"`
}

var poolReassignStoreRequest = sync.Pool{
	New: func() any {
		return new(ReassignStoreRequest)
	},
}

// GetReassignStoreRequest() 从对象池中获取ReassignStoreRequest
func GetReassignStoreRequest() *ReassignStoreRequest {
	return poolReassignStoreRequest.Get().(*ReassignStoreRequest)
}

// ReleaseReassignStoreRequest 释放ReassignStoreRequest
func ReleaseReassignStoreRequest(v *ReassignStoreRequest) {
	v.WorkcardId = 0
	v.TargetServiceStoreId = 0
	v.OriginServiceStoreId = 0
	poolReassignStoreRequest.Put(v)
}
