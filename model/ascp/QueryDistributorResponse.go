package ascp

import (
	"sync"
)

// QueryDistributorResponse 结构体
type QueryDistributorResponse struct {
	// 分销商列表
	MerchantList []Integer `json:"merchant_list,omitempty" xml:"merchant_list>integer,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 每页条数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

var poolQueryDistributorResponse = sync.Pool{
	New: func() any {
		return new(QueryDistributorResponse)
	},
}

// GetQueryDistributorResponse() 从对象池中获取QueryDistributorResponse
func GetQueryDistributorResponse() *QueryDistributorResponse {
	return poolQueryDistributorResponse.Get().(*QueryDistributorResponse)
}

// ReleaseQueryDistributorResponse 释放QueryDistributorResponse
func ReleaseQueryDistributorResponse(v *QueryDistributorResponse) {
	v.MerchantList = v.MerchantList[:0]
	v.CurrentPage = 0
	v.PageSize = 0
	v.TotalPage = 0
	poolQueryDistributorResponse.Put(v)
}
