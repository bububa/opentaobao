package simba

import (
	"sync"
)

// TopAccountStatusVo 结构体
type TopAccountStatusVo struct {
	// 是否是新bp用户
	IsUniversalBpUser bool `json:"is_universal_bp_user,omitempty" xml:"is_universal_bp_user,omitempty"`
}

var poolTopAccountStatusVo = sync.Pool{
	New: func() any {
		return new(TopAccountStatusVo)
	},
}

// GetTopAccountStatusVo() 从对象池中获取TopAccountStatusVo
func GetTopAccountStatusVo() *TopAccountStatusVo {
	return poolTopAccountStatusVo.Get().(*TopAccountStatusVo)
}

// ReleaseTopAccountStatusVo 释放TopAccountStatusVo
func ReleaseTopAccountStatusVo(v *TopAccountStatusVo) {
	v.IsUniversalBpUser = false
	poolTopAccountStatusVo.Put(v)
}
