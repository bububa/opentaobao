package ascp

import (
	"sync"
)

// BatchCreateItemMappingRequest 结构体
type BatchCreateItemMappingRequest struct {
	// 商货品关联列表，批量数量不可超过30
	ItemMappings []ItemMapping `json:"item_mappings,omitempty" xml:"item_mappings>item_mapping,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolBatchCreateItemMappingRequest = sync.Pool{
	New: func() any {
		return new(BatchCreateItemMappingRequest)
	},
}

// GetBatchCreateItemMappingRequest() 从对象池中获取BatchCreateItemMappingRequest
func GetBatchCreateItemMappingRequest() *BatchCreateItemMappingRequest {
	return poolBatchCreateItemMappingRequest.Get().(*BatchCreateItemMappingRequest)
}

// ReleaseBatchCreateItemMappingRequest 释放BatchCreateItemMappingRequest
func ReleaseBatchCreateItemMappingRequest(v *BatchCreateItemMappingRequest) {
	v.ItemMappings = v.ItemMappings[:0]
	v.RequestId = ""
	v.RequestTime = 0
	poolBatchCreateItemMappingRequest.Put(v)
}
