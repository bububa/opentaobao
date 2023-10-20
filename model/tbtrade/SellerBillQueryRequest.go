package tbtrade

import (
	"sync"
)

// SellerBillQueryRequest 结构体
type SellerBillQueryRequest struct {
	// 账单时间
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 页号，从1开始
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小，不得超过1000
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolSellerBillQueryRequest = sync.Pool{
	New: func() any {
		return new(SellerBillQueryRequest)
	},
}

// GetSellerBillQueryRequest() 从对象池中获取SellerBillQueryRequest
func GetSellerBillQueryRequest() *SellerBillQueryRequest {
	return poolSellerBillQueryRequest.Get().(*SellerBillQueryRequest)
}

// ReleaseSellerBillQueryRequest 释放SellerBillQueryRequest
func ReleaseSellerBillQueryRequest(v *SellerBillQueryRequest) {
	v.BizDate = ""
	v.PageNo = 0
	v.PageSize = 0
	poolSellerBillQueryRequest.Put(v)
}
