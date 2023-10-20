package security

import (
	"sync"
)

// RpInitResultBo 结构体
type RpInitResultBo struct {
	// token
	VerifyToken string `json:"verify_token,omitempty" xml:"verify_token,omitempty"`
	// 过期时间
	Expire int64 `json:"expire,omitempty" xml:"expire,omitempty"`
}

var poolRpInitResultBo = sync.Pool{
	New: func() any {
		return new(RpInitResultBo)
	},
}

// GetRpInitResultBo() 从对象池中获取RpInitResultBo
func GetRpInitResultBo() *RpInitResultBo {
	return poolRpInitResultBo.Get().(*RpInitResultBo)
}

// ReleaseRpInitResultBo 释放RpInitResultBo
func ReleaseRpInitResultBo(v *RpInitResultBo) {
	v.VerifyToken = ""
	v.Expire = 0
	poolRpInitResultBo.Put(v)
}
