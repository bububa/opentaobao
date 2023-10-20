package alitripmerchant

import (
	"sync"
)

// BaseDescVo 结构体
type BaseDescVo struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
}

var poolBaseDescVo = sync.Pool{
	New: func() any {
		return new(BaseDescVo)
	},
}

// GetBaseDescVo() 从对象池中获取BaseDescVo
func GetBaseDescVo() *BaseDescVo {
	return poolBaseDescVo.Get().(*BaseDescVo)
}

// ReleaseBaseDescVo 释放BaseDescVo
func ReleaseBaseDescVo(v *BaseDescVo) {
	v.Name = ""
	v.Description = ""
	poolBaseDescVo.Put(v)
}
