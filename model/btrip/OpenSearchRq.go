package btrip

import (
	"sync"
)

// OpenSearchRq 结构体
type OpenSearchRq struct {
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 部门id
	DepartId string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 更新时间大于等于此时间的审批单
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 更新开始时间
	UpdateEndTime string `json:"update_end_time,omitempty" xml:"update_end_time,omitempty"`
	// 更新结束时间
	UpdateStartTime string `json:"update_start_time,omitempty" xml:"update_start_time,omitempty"`
	// 页数，从1开始
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页返回数量，默认10，最多50
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 申请单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
}

var poolOpenSearchRq = sync.Pool{
	New: func() any {
		return new(OpenSearchRq)
	},
}

// GetOpenSearchRq() 从对象池中获取OpenSearchRq
func GetOpenSearchRq() *OpenSearchRq {
	return poolOpenSearchRq.Get().(*OpenSearchRq)
}

// ReleaseOpenSearchRq 释放OpenSearchRq
func ReleaseOpenSearchRq(v *OpenSearchRq) {
	v.StartTime = ""
	v.DepartId = ""
	v.EndTime = ""
	v.CorpId = ""
	v.UserId = ""
	v.GmtModified = ""
	v.UpdateEndTime = ""
	v.UpdateStartTime = ""
	v.Page = 0
	v.PageSize = 0
	v.ApplyId = 0
	poolOpenSearchRq.Put(v)
}
