package ascp

import (
	"sync"
)

// ItemBatchUpdateAsyncRequest 结构体
type ItemBatchUpdateAsyncRequest struct {
	// 货品列表
	ScItems []ScItem `json:"sc_items,omitempty" xml:"sc_items>sc_item,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolItemBatchUpdateAsyncRequest = sync.Pool{
	New: func() any {
		return new(ItemBatchUpdateAsyncRequest)
	},
}

// GetItemBatchUpdateAsyncRequest() 从对象池中获取ItemBatchUpdateAsyncRequest
func GetItemBatchUpdateAsyncRequest() *ItemBatchUpdateAsyncRequest {
	return poolItemBatchUpdateAsyncRequest.Get().(*ItemBatchUpdateAsyncRequest)
}

// ReleaseItemBatchUpdateAsyncRequest 释放ItemBatchUpdateAsyncRequest
func ReleaseItemBatchUpdateAsyncRequest(v *ItemBatchUpdateAsyncRequest) {
	v.ScItems = v.ScItems[:0]
	v.RequestId = ""
	v.RequestTime = 0
	poolItemBatchUpdateAsyncRequest.Put(v)
}
