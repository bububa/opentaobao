package campus

import (
	"sync"
)

// GetResourceBookInfoRequest 结构体
type GetResourceBookInfoRequest struct {
	// 申请结束时间
	ApplyEndTime string `json:"apply_end_time,omitempty" xml:"apply_end_time,omitempty"`
	// 申请开始时间
	ApplyStartTime string `json:"apply_start_time,omitempty" xml:"apply_start_time,omitempty"`
	// 资源类型
	ResourceTypeCode string `json:"resource_type_code,omitempty" xml:"resource_type_code,omitempty"`
	// 一页大小
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 园区id
	CampusId int64 `json:"campus_id,omitempty" xml:"campus_id,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
}

var poolGetResourceBookInfoRequest = sync.Pool{
	New: func() any {
		return new(GetResourceBookInfoRequest)
	},
}

// GetGetResourceBookInfoRequest() 从对象池中获取GetResourceBookInfoRequest
func GetGetResourceBookInfoRequest() *GetResourceBookInfoRequest {
	return poolGetResourceBookInfoRequest.Get().(*GetResourceBookInfoRequest)
}

// ReleaseGetResourceBookInfoRequest 释放GetResourceBookInfoRequest
func ReleaseGetResourceBookInfoRequest(v *GetResourceBookInfoRequest) {
	v.ApplyEndTime = ""
	v.ApplyStartTime = ""
	v.ResourceTypeCode = ""
	v.Size = 0
	v.CampusId = 0
	v.Page = 0
	poolGetResourceBookInfoRequest.Put(v)
}
