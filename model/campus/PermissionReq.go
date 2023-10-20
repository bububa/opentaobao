package campus

import (
	"sync"
)

// PermissionReq 结构体
type PermissionReq struct {
	// 权限key
	PrivKey string `json:"priv_key,omitempty" xml:"priv_key,omitempty"`
	// 生效时间
	EffectiveTime string `json:"effective_time,omitempty" xml:"effective_time,omitempty"`
	// 过期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
}

var poolPermissionReq = sync.Pool{
	New: func() any {
		return new(PermissionReq)
	},
}

// GetPermissionReq() 从对象池中获取PermissionReq
func GetPermissionReq() *PermissionReq {
	return poolPermissionReq.Get().(*PermissionReq)
}

// ReleasePermissionReq 释放PermissionReq
func ReleasePermissionReq(v *PermissionReq) {
	v.PrivKey = ""
	v.EffectiveTime = ""
	v.ExpireTime = ""
	poolPermissionReq.Put(v)
}
