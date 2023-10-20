package ascp

import (
	"sync"
)

// BatchCreateScItemRequest 结构体
type BatchCreateScItemRequest struct {
	// 货品列表，批量数量不大于30
	ScItems []ScItem `json:"sc_items,omitempty" xml:"sc_items>sc_item,omitempty"`
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolBatchCreateScItemRequest = sync.Pool{
	New: func() any {
		return new(BatchCreateScItemRequest)
	},
}

// GetBatchCreateScItemRequest() 从对象池中获取BatchCreateScItemRequest
func GetBatchCreateScItemRequest() *BatchCreateScItemRequest {
	return poolBatchCreateScItemRequest.Get().(*BatchCreateScItemRequest)
}

// ReleaseBatchCreateScItemRequest 释放BatchCreateScItemRequest
func ReleaseBatchCreateScItemRequest(v *BatchCreateScItemRequest) {
	v.ScItems = v.ScItems[:0]
	v.RequestId = ""
	v.RequestTime = 0
	poolBatchCreateScItemRequest.Put(v)
}
