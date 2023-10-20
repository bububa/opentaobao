package alihealthpw

import (
	"sync"
)

// PendingListReq 结构体
type PendingListReq struct {
	// 状态列表
	ApplyStatusList []string `json:"apply_status_list,omitempty" xml:"apply_status_list>string,omitempty"`
	// 申请类型
	ApplyType string `json:"apply_type,omitempty" xml:"apply_type,omitempty"`
	// 当前页面
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPendingListReq = sync.Pool{
	New: func() any {
		return new(PendingListReq)
	},
}

// GetPendingListReq() 从对象池中获取PendingListReq
func GetPendingListReq() *PendingListReq {
	return poolPendingListReq.Get().(*PendingListReq)
}

// ReleasePendingListReq 释放PendingListReq
func ReleasePendingListReq(v *PendingListReq) {
	v.ApplyStatusList = v.ApplyStatusList[:0]
	v.ApplyType = ""
	v.PageNumber = 0
	v.PageSize = 0
	poolPendingListReq.Put(v)
}
