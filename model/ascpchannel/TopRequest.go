package ascpchannel

import (
	"sync"
)

// TopRequest 结构体
type TopRequest struct {
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 创建供品关系开始时间
	BeginTime string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 创建供品关系结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 每页条数，最大不超过200
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}

var poolTopRequest = sync.Pool{
	New: func() any {
		return new(TopRequest)
	},
}

// GetTopRequest() 从对象池中获取TopRequest
func GetTopRequest() *TopRequest {
	return poolTopRequest.Get().(*TopRequest)
}

// ReleaseTopRequest 释放TopRequest
func ReleaseTopRequest(v *TopRequest) {
	v.SupplierId = ""
	v.BeginTime = ""
	v.EndTime = ""
	v.PageSize = 0
	v.CurrentPage = 0
	poolTopRequest.Put(v)
}
