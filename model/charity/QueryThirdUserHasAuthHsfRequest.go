package charity

import (
	"sync"
)

// QueryThirdUserHasAuthHsfRequest 结构体
type QueryThirdUserHasAuthHsfRequest struct {
	// appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 用户唯一标识
	UserKey string `json:"user_key,omitempty" xml:"user_key,omitempty"`
}

var poolQueryThirdUserHasAuthHsfRequest = sync.Pool{
	New: func() any {
		return new(QueryThirdUserHasAuthHsfRequest)
	},
}

// GetQueryThirdUserHasAuthHsfRequest() 从对象池中获取QueryThirdUserHasAuthHsfRequest
func GetQueryThirdUserHasAuthHsfRequest() *QueryThirdUserHasAuthHsfRequest {
	return poolQueryThirdUserHasAuthHsfRequest.Get().(*QueryThirdUserHasAuthHsfRequest)
}

// ReleaseQueryThirdUserHasAuthHsfRequest 释放QueryThirdUserHasAuthHsfRequest
func ReleaseQueryThirdUserHasAuthHsfRequest(v *QueryThirdUserHasAuthHsfRequest) {
	v.AppKey = ""
	v.UserKey = ""
	poolQueryThirdUserHasAuthHsfRequest.Put(v)
}
