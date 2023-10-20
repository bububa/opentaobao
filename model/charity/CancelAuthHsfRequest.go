package charity

import (
	"sync"
)

// CancelAuthHsfRequest 结构体
type CancelAuthHsfRequest struct {
	// appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 用户唯一标识
	UserKey string `json:"user_key,omitempty" xml:"user_key,omitempty"`
}

var poolCancelAuthHsfRequest = sync.Pool{
	New: func() any {
		return new(CancelAuthHsfRequest)
	},
}

// GetCancelAuthHsfRequest() 从对象池中获取CancelAuthHsfRequest
func GetCancelAuthHsfRequest() *CancelAuthHsfRequest {
	return poolCancelAuthHsfRequest.Get().(*CancelAuthHsfRequest)
}

// ReleaseCancelAuthHsfRequest 释放CancelAuthHsfRequest
func ReleaseCancelAuthHsfRequest(v *CancelAuthHsfRequest) {
	v.AppKey = ""
	v.UserKey = ""
	poolCancelAuthHsfRequest.Put(v)
}
