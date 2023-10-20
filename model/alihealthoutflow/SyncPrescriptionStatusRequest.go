package alihealthoutflow

import (
	"sync"
)

// SyncPrescriptionStatusRequest 结构体
type SyncPrescriptionStatusRequest struct {
	// 阿里将康处方号(非空)
	RxNo string `json:"rx_no,omitempty" xml:"rx_no,omitempty"`
	// 处方状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolSyncPrescriptionStatusRequest = sync.Pool{
	New: func() any {
		return new(SyncPrescriptionStatusRequest)
	},
}

// GetSyncPrescriptionStatusRequest() 从对象池中获取SyncPrescriptionStatusRequest
func GetSyncPrescriptionStatusRequest() *SyncPrescriptionStatusRequest {
	return poolSyncPrescriptionStatusRequest.Get().(*SyncPrescriptionStatusRequest)
}

// ReleaseSyncPrescriptionStatusRequest 释放SyncPrescriptionStatusRequest
func ReleaseSyncPrescriptionStatusRequest(v *SyncPrescriptionStatusRequest) {
	v.RxNo = ""
	v.Status = ""
	poolSyncPrescriptionStatusRequest.Put(v)
}
