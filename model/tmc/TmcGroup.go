package tmc

import (
	"sync"
)

// TmcGroup 结构体
type TmcGroup struct {
	// 分组名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolTmcGroup = sync.Pool{
	New: func() any {
		return new(TmcGroup)
	},
}

// GetTmcGroup() 从对象池中获取TmcGroup
func GetTmcGroup() *TmcGroup {
	return poolTmcGroup.Get().(*TmcGroup)
}

// ReleaseTmcGroup 释放TmcGroup
func ReleaseTmcGroup(v *TmcGroup) {
	v.Name = ""
	poolTmcGroup.Put(v)
}
