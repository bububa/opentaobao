package wdk

import (
	"sync"
)

// ItemPoolActivityQueryRequest 结构体
type ItemPoolActivityQueryRequest struct {
	// 外部活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 活动ID
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
}

var poolItemPoolActivityQueryRequest = sync.Pool{
	New: func() any {
		return new(ItemPoolActivityQueryRequest)
	},
}

// GetItemPoolActivityQueryRequest() 从对象池中获取ItemPoolActivityQueryRequest
func GetItemPoolActivityQueryRequest() *ItemPoolActivityQueryRequest {
	return poolItemPoolActivityQueryRequest.Get().(*ItemPoolActivityQueryRequest)
}

// ReleaseItemPoolActivityQueryRequest 释放ItemPoolActivityQueryRequest
func ReleaseItemPoolActivityQueryRequest(v *ItemPoolActivityQueryRequest) {
	v.OutActId = ""
	v.ActId = 0
	poolItemPoolActivityQueryRequest.Put(v)
}
