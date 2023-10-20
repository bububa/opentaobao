package wlbimports

import (
	"sync"
)

// TransferStoreResponse 结构体
type TransferStoreResponse struct {
	// 集货仓信息列表
	TransferStoreInfoResponseList []TransferStoreInfoResponse `json:"transfer_store_info_response_list,omitempty" xml:"transfer_store_info_response_list>transfer_store_info_response,omitempty"`
}

var poolTransferStoreResponse = sync.Pool{
	New: func() any {
		return new(TransferStoreResponse)
	},
}

// GetTransferStoreResponse() 从对象池中获取TransferStoreResponse
func GetTransferStoreResponse() *TransferStoreResponse {
	return poolTransferStoreResponse.Get().(*TransferStoreResponse)
}

// ReleaseTransferStoreResponse 释放TransferStoreResponse
func ReleaseTransferStoreResponse(v *TransferStoreResponse) {
	v.TransferStoreInfoResponseList = v.TransferStoreInfoResponseList[:0]
	poolTransferStoreResponse.Put(v)
}
