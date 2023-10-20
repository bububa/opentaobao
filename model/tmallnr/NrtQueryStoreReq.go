package tmallnr

import (
	"sync"
)

// NrtQueryStoreReq 结构体
type NrtQueryStoreReq struct {
	// 活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 每页数据条数，最大100条
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
}

var poolNrtQueryStoreReq = sync.Pool{
	New: func() any {
		return new(NrtQueryStoreReq)
	},
}

// GetNrtQueryStoreReq() 从对象池中获取NrtQueryStoreReq
func GetNrtQueryStoreReq() *NrtQueryStoreReq {
	return poolNrtQueryStoreReq.Get().(*NrtQueryStoreReq)
}

// ReleaseNrtQueryStoreReq 释放NrtQueryStoreReq
func ReleaseNrtQueryStoreReq(v *NrtQueryStoreReq) {
	v.ActivityId = 0
	v.PageSize = 0
	v.PageIndex = 0
	poolNrtQueryStoreReq.Put(v)
}
