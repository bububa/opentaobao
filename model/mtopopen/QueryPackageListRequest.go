package mtopopen

import (
	"sync"
)

// QueryPackageListRequest 结构体
type QueryPackageListRequest struct {
	// 快递公司编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 用户的唯一标识
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 页数，从1开始
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolQueryPackageListRequest = sync.Pool{
	New: func() any {
		return new(QueryPackageListRequest)
	},
}

// GetQueryPackageListRequest() 从对象池中获取QueryPackageListRequest
func GetQueryPackageListRequest() *QueryPackageListRequest {
	return poolQueryPackageListRequest.Get().(*QueryPackageListRequest)
}

// ReleaseQueryPackageListRequest 释放QueryPackageListRequest
func ReleaseQueryPackageListRequest(v *QueryPackageListRequest) {
	v.CpCode = ""
	v.OpenId = ""
	v.PageNo = 0
	v.PageSize = 0
	poolQueryPackageListRequest.Put(v)
}
