package promotion

import (
	"sync"
)

// UseResultDo 结构体
type UseResultDo struct {
	// isUsed
	IsUsed bool `json:"is_used,omitempty" xml:"is_used,omitempty"`
}

var poolUseResultDo = sync.Pool{
	New: func() any {
		return new(UseResultDo)
	},
}

// GetUseResultDo() 从对象池中获取UseResultDo
func GetUseResultDo() *UseResultDo {
	return poolUseResultDo.Get().(*UseResultDo)
}

// ReleaseUseResultDo 释放UseResultDo
func ReleaseUseResultDo(v *UseResultDo) {
	v.IsUsed = false
	poolUseResultDo.Put(v)
}
