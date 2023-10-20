package tmallnr

import (
	"sync"
)

// NrtQueryActivityReq 结构体
type NrtQueryActivityReq struct {
	// 活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

var poolNrtQueryActivityReq = sync.Pool{
	New: func() any {
		return new(NrtQueryActivityReq)
	},
}

// GetNrtQueryActivityReq() 从对象池中获取NrtQueryActivityReq
func GetNrtQueryActivityReq() *NrtQueryActivityReq {
	return poolNrtQueryActivityReq.Get().(*NrtQueryActivityReq)
}

// ReleaseNrtQueryActivityReq 释放NrtQueryActivityReq
func ReleaseNrtQueryActivityReq(v *NrtQueryActivityReq) {
	v.ActivityId = 0
	poolNrtQueryActivityReq.Put(v)
}
