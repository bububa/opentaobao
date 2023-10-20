package wlbimports

import (
	"sync"
)

// TransferStoreInfoResponse 结构体
type TransferStoreInfoResponse struct {
	// 仓库揽收信息
	PickUpResponse *PickupInfo `json:"pick_up_response,omitempty" xml:"pick_up_response,omitempty"`
	// 集货仓信息
	StoreCenterResponse *TransferStoreCenterResponse `json:"store_center_response,omitempty" xml:"store_center_response,omitempty"`
}

var poolTransferStoreInfoResponse = sync.Pool{
	New: func() any {
		return new(TransferStoreInfoResponse)
	},
}

// GetTransferStoreInfoResponse() 从对象池中获取TransferStoreInfoResponse
func GetTransferStoreInfoResponse() *TransferStoreInfoResponse {
	return poolTransferStoreInfoResponse.Get().(*TransferStoreInfoResponse)
}

// ReleaseTransferStoreInfoResponse 释放TransferStoreInfoResponse
func ReleaseTransferStoreInfoResponse(v *TransferStoreInfoResponse) {
	v.PickUpResponse = nil
	v.StoreCenterResponse = nil
	poolTransferStoreInfoResponse.Put(v)
}
