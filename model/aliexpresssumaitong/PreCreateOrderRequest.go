package aliexpresssumaitong

import (
	"sync"
)

// PreCreateOrderRequest 结构体
type PreCreateOrderRequest struct {
	// 商品信息
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}

var poolPreCreateOrderRequest = sync.Pool{
	New: func() any {
		return new(PreCreateOrderRequest)
	},
}

// GetPreCreateOrderRequest() 从对象池中获取PreCreateOrderRequest
func GetPreCreateOrderRequest() *PreCreateOrderRequest {
	return poolPreCreateOrderRequest.Get().(*PreCreateOrderRequest)
}

// ReleasePreCreateOrderRequest 释放PreCreateOrderRequest
func ReleasePreCreateOrderRequest(v *PreCreateOrderRequest) {
	v.Items = v.Items[:0]
	poolPreCreateOrderRequest.Put(v)
}
