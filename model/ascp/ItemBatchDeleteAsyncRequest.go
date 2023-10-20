package ascp

import (
	"sync"
)

// ItemBatchDeleteAsyncRequest 结构体
type ItemBatchDeleteAsyncRequest struct {
	// 货品列表
	ScItems []ScItem `json:"sc_items,omitempty" xml:"sc_items>sc_item,omitempty"`
	// 业务请求ID，用于做幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ERP调翱象接口创建货品的时间戳
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
}

var poolItemBatchDeleteAsyncRequest = sync.Pool{
	New: func() any {
		return new(ItemBatchDeleteAsyncRequest)
	},
}

// GetItemBatchDeleteAsyncRequest() 从对象池中获取ItemBatchDeleteAsyncRequest
func GetItemBatchDeleteAsyncRequest() *ItemBatchDeleteAsyncRequest {
	return poolItemBatchDeleteAsyncRequest.Get().(*ItemBatchDeleteAsyncRequest)
}

// ReleaseItemBatchDeleteAsyncRequest 释放ItemBatchDeleteAsyncRequest
func ReleaseItemBatchDeleteAsyncRequest(v *ItemBatchDeleteAsyncRequest) {
	v.ScItems = v.ScItems[:0]
	v.RequestId = ""
	v.RequestTime = 0
	poolItemBatchDeleteAsyncRequest.Put(v)
}
