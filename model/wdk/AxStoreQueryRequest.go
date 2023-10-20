package wdk

import (
	"sync"
)

// AxStoreQueryRequest 结构体
type AxStoreQueryRequest struct {
	// 经营店code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
}

var poolAxStoreQueryRequest = sync.Pool{
	New: func() any {
		return new(AxStoreQueryRequest)
	},
}

// GetAxStoreQueryRequest() 从对象池中获取AxStoreQueryRequest
func GetAxStoreQueryRequest() *AxStoreQueryRequest {
	return poolAxStoreQueryRequest.Get().(*AxStoreQueryRequest)
}

// ReleaseAxStoreQueryRequest 释放AxStoreQueryRequest
func ReleaseAxStoreQueryRequest(v *AxStoreQueryRequest) {
	v.StoreCode = ""
	poolAxStoreQueryRequest.Put(v)
}
