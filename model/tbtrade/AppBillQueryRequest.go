package tbtrade

import (
	"sync"
)

// AppBillQueryRequest 结构体
type AppBillQueryRequest struct {
	// 账单时间
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 页号，从1开始
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小，不得超过1000
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolAppBillQueryRequest = sync.Pool{
	New: func() any {
		return new(AppBillQueryRequest)
	},
}

// GetAppBillQueryRequest() 从对象池中获取AppBillQueryRequest
func GetAppBillQueryRequest() *AppBillQueryRequest {
	return poolAppBillQueryRequest.Get().(*AppBillQueryRequest)
}

// ReleaseAppBillQueryRequest 释放AppBillQueryRequest
func ReleaseAppBillQueryRequest(v *AppBillQueryRequest) {
	v.BizDate = ""
	v.PageNo = 0
	v.PageSize = 0
	poolAppBillQueryRequest.Put(v)
}
