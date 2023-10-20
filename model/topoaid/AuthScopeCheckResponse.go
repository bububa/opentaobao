package topoaid

import (
	"sync"
)

// AuthScopeCheckResponse 结构体
type AuthScopeCheckResponse struct {
	// 授权列表
	ScopeNames []string `json:"scope_names,omitempty" xml:"scope_names>string,omitempty"`
	// 授权到期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
}

var poolAuthScopeCheckResponse = sync.Pool{
	New: func() any {
		return new(AuthScopeCheckResponse)
	},
}

// GetAuthScopeCheckResponse() 从对象池中获取AuthScopeCheckResponse
func GetAuthScopeCheckResponse() *AuthScopeCheckResponse {
	return poolAuthScopeCheckResponse.Get().(*AuthScopeCheckResponse)
}

// ReleaseAuthScopeCheckResponse 释放AuthScopeCheckResponse
func ReleaseAuthScopeCheckResponse(v *AuthScopeCheckResponse) {
	v.ScopeNames = v.ScopeNames[:0]
	v.ExpireTime = ""
	poolAuthScopeCheckResponse.Put(v)
}
