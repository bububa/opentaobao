package axindata

import (
	"sync"
)

// PolicyVo 结构体
type PolicyVo struct {
	// 扣除值
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// 往前推小时
	Hour int64 `json:"hour,omitempty" xml:"hour,omitempty"`
}

var poolPolicyVo = sync.Pool{
	New: func() any {
		return new(PolicyVo)
	},
}

// GetPolicyVo() 从对象池中获取PolicyVo
func GetPolicyVo() *PolicyVo {
	return poolPolicyVo.Get().(*PolicyVo)
}

// ReleasePolicyVo 释放PolicyVo
func ReleasePolicyVo(v *PolicyVo) {
	v.Value = 0
	v.Hour = 0
	poolPolicyVo.Put(v)
}
