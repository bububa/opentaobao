package alimember

import (
	"sync"
)

// PageQueryIsvCustomerRequest 结构体
type PageQueryIsvCustomerRequest struct {
	// 商户id
	OpenMerchantId string `json:"open_merchant_id,omitempty" xml:"open_merchant_id,omitempty"`
	// 第几页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页大小，最大不超过1000
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPageQueryIsvCustomerRequest = sync.Pool{
	New: func() any {
		return new(PageQueryIsvCustomerRequest)
	},
}

// GetPageQueryIsvCustomerRequest() 从对象池中获取PageQueryIsvCustomerRequest
func GetPageQueryIsvCustomerRequest() *PageQueryIsvCustomerRequest {
	return poolPageQueryIsvCustomerRequest.Get().(*PageQueryIsvCustomerRequest)
}

// ReleasePageQueryIsvCustomerRequest 释放PageQueryIsvCustomerRequest
func ReleasePageQueryIsvCustomerRequest(v *PageQueryIsvCustomerRequest) {
	v.OpenMerchantId = ""
	v.PageNo = 0
	v.PageSize = 0
	poolPageQueryIsvCustomerRequest.Put(v)
}
