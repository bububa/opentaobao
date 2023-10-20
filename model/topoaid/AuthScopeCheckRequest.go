package topoaid

import (
	"sync"
)

// AuthScopeCheckRequest 结构体
type AuthScopeCheckRequest struct {
	// 用户手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 从context中获取的用户标识
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
}

var poolAuthScopeCheckRequest = sync.Pool{
	New: func() any {
		return new(AuthScopeCheckRequest)
	},
}

// GetAuthScopeCheckRequest() 从对象池中获取AuthScopeCheckRequest
func GetAuthScopeCheckRequest() *AuthScopeCheckRequest {
	return poolAuthScopeCheckRequest.Get().(*AuthScopeCheckRequest)
}

// ReleaseAuthScopeCheckRequest 释放AuthScopeCheckRequest
func ReleaseAuthScopeCheckRequest(v *AuthScopeCheckRequest) {
	v.Mobile = ""
	v.OpenId = ""
	poolAuthScopeCheckRequest.Put(v)
}
