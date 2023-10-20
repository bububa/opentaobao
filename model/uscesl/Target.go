package uscesl

import (
	"sync"
)

// Target 结构体
type Target struct {
	// AP的mac地址
	Mac string `json:"mac,omitempty" xml:"mac,omitempty"`
	// 型号
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 是否在线
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
	// 是否激活
	IsActivate bool `json:"is_activate,omitempty" xml:"is_activate,omitempty"`
}

var poolTarget = sync.Pool{
	New: func() any {
		return new(Target)
	},
}

// GetTarget() 从对象池中获取Target
func GetTarget() *Target {
	return poolTarget.Get().(*Target)
}

// ReleaseTarget 释放Target
func ReleaseTarget(v *Target) {
	v.Mac = ""
	v.Model = ""
	v.Status = false
	v.IsActivate = false
	poolTarget.Put(v)
}
