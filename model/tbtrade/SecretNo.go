package tbtrade

import (
	"sync"
)

// SecretNo 结构体
type SecretNo struct {
	// 隐私号
	SecretNo string `json:"secret_no,omitempty" xml:"secret_no,omitempty"`
	// 延期后隐私号的过期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
}

var poolSecretNo = sync.Pool{
	New: func() any {
		return new(SecretNo)
	},
}

// GetSecretNo() 从对象池中获取SecretNo
func GetSecretNo() *SecretNo {
	return poolSecretNo.Get().(*SecretNo)
}

// ReleaseSecretNo 释放SecretNo
func ReleaseSecretNo(v *SecretNo) {
	v.SecretNo = ""
	v.ExpireTime = ""
	poolSecretNo.Put(v)
}
