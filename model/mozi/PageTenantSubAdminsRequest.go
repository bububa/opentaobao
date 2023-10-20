package mozi

import (
	"sync"
)

// PageTenantSubAdminsRequest 结构体
type PageTenantSubAdminsRequest struct {
	// 页数
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 租户id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
	// 一页的条目
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否返回总数量
	ReturnTotalSize bool `json:"return_total_size,omitempty" xml:"return_total_size,omitempty"`
}

var poolPageTenantSubAdminsRequest = sync.Pool{
	New: func() any {
		return new(PageTenantSubAdminsRequest)
	},
}

// GetPageTenantSubAdminsRequest() 从对象池中获取PageTenantSubAdminsRequest
func GetPageTenantSubAdminsRequest() *PageTenantSubAdminsRequest {
	return poolPageTenantSubAdminsRequest.Get().(*PageTenantSubAdminsRequest)
}

// ReleasePageTenantSubAdminsRequest 释放PageTenantSubAdminsRequest
func ReleasePageTenantSubAdminsRequest(v *PageTenantSubAdminsRequest) {
	v.PageNo = 0
	v.TenantId = 0
	v.PageSize = 0
	v.ReturnTotalSize = false
	poolPageTenantSubAdminsRequest.Put(v)
}
