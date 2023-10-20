package kclub

import (
	"sync"
)

// TenancyAuth 结构体
type TenancyAuth struct {
	// 鉴权秘钥
	SecretKey string `json:"secret_key,omitempty" xml:"secret_key,omitempty"`
	// 鉴权应用名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolTenancyAuth = sync.Pool{
	New: func() any {
		return new(TenancyAuth)
	},
}

// GetTenancyAuth() 从对象池中获取TenancyAuth
func GetTenancyAuth() *TenancyAuth {
	return poolTenancyAuth.Get().(*TenancyAuth)
}

// ReleaseTenancyAuth 释放TenancyAuth
func ReleaseTenancyAuth(v *TenancyAuth) {
	v.SecretKey = ""
	v.Name = ""
	poolTenancyAuth.Put(v)
}
