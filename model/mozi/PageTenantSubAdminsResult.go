package mozi

import (
	"sync"
)

// PageTenantSubAdminsResult 结构体
type PageTenantSubAdminsResult struct {
	// data数据
	Datas []TenantAdmin `json:"datas,omitempty" xml:"datas>tenant_admin,omitempty"`
	// requestId
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// responseMessage
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// responseCode
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 总条数
	TotalSize int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
	// 页的条目数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页数
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPageTenantSubAdminsResult = sync.Pool{
	New: func() any {
		return new(PageTenantSubAdminsResult)
	},
}

// GetPageTenantSubAdminsResult() 从对象池中获取PageTenantSubAdminsResult
func GetPageTenantSubAdminsResult() *PageTenantSubAdminsResult {
	return poolPageTenantSubAdminsResult.Get().(*PageTenantSubAdminsResult)
}

// ReleasePageTenantSubAdminsResult 释放PageTenantSubAdminsResult
func ReleasePageTenantSubAdminsResult(v *PageTenantSubAdminsResult) {
	v.Datas = v.Datas[:0]
	v.RequestId = ""
	v.ResponseMessage = ""
	v.ResponseCode = ""
	v.TotalSize = 0
	v.PageSize = 0
	v.CurrentPage = 0
	v.Success = false
	poolPageTenantSubAdminsResult.Put(v)
}
