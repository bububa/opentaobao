package wdk

import (
	"sync"
)

// CreateFetchResponse 结构体
type CreateFetchResponse struct {
	// 取货单di
	ReferId string `json:"refer_id,omitempty" xml:"refer_id,omitempty"`
}

var poolCreateFetchResponse = sync.Pool{
	New: func() any {
		return new(CreateFetchResponse)
	},
}

// GetCreateFetchResponse() 从对象池中获取CreateFetchResponse
func GetCreateFetchResponse() *CreateFetchResponse {
	return poolCreateFetchResponse.Get().(*CreateFetchResponse)
}

// ReleaseCreateFetchResponse 释放CreateFetchResponse
func ReleaseCreateFetchResponse(v *CreateFetchResponse) {
	v.ReferId = ""
	poolCreateFetchResponse.Put(v)
}
