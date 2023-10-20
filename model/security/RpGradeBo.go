package security

import (
	"sync"
)

// RpGradeBo 结构体
type RpGradeBo struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// level
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
}

var poolRpGradeBo = sync.Pool{
	New: func() any {
		return new(RpGradeBo)
	},
}

// GetRpGradeBo() 从对象池中获取RpGradeBo
func GetRpGradeBo() *RpGradeBo {
	return poolRpGradeBo.Get().(*RpGradeBo)
}

// ReleaseRpGradeBo 释放RpGradeBo
func ReleaseRpGradeBo(v *RpGradeBo) {
	v.Desc = ""
	v.Name = ""
	v.Level = 0
	poolRpGradeBo.Put(v)
}
