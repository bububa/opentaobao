package ascp

import (
	"sync"
)

// QueryDistributorRequest 结构体
type QueryDistributorRequest struct {
	// 业务请求ID，用于幂等
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 业务请求时间。时间戳。 毫秒
	RequestTime int64 `json:"request_time,omitempty" xml:"request_time,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 一页多少条
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolQueryDistributorRequest = sync.Pool{
	New: func() any {
		return new(QueryDistributorRequest)
	},
}

// GetQueryDistributorRequest() 从对象池中获取QueryDistributorRequest
func GetQueryDistributorRequest() *QueryDistributorRequest {
	return poolQueryDistributorRequest.Get().(*QueryDistributorRequest)
}

// ReleaseQueryDistributorRequest 释放QueryDistributorRequest
func ReleaseQueryDistributorRequest(v *QueryDistributorRequest) {
	v.RequestId = ""
	v.RequestTime = 0
	v.CurrentPage = 0
	v.PageSize = 0
	poolQueryDistributorRequest.Put(v)
}
