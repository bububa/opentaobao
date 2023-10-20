package jstinteractive

import (
	"sync"
)

// InteractivePointQueryResponse 结构体
type InteractivePointQueryResponse struct {
	// 用户积分总额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolInteractivePointQueryResponse = sync.Pool{
	New: func() any {
		return new(InteractivePointQueryResponse)
	},
}

// GetInteractivePointQueryResponse() 从对象池中获取InteractivePointQueryResponse
func GetInteractivePointQueryResponse() *InteractivePointQueryResponse {
	return poolInteractivePointQueryResponse.Get().(*InteractivePointQueryResponse)
}

// ReleaseInteractivePointQueryResponse 释放InteractivePointQueryResponse
func ReleaseInteractivePointQueryResponse(v *InteractivePointQueryResponse) {
	v.Amount = 0
	poolInteractivePointQueryResponse.Put(v)
}
