package ott

import (
	"sync"
)

// LauncherDo 结构体
type LauncherDo struct {
	// 桌面坑位
	Items []Items `json:"items,omitempty" xml:"items>items,omitempty"`
	// 设备属性
	Property *PropertyDo `json:"property,omitempty" xml:"property,omitempty"`
	// 桌面配置
	Version *VersionDo `json:"version,omitempty" xml:"version,omitempty"`
}

var poolLauncherDo = sync.Pool{
	New: func() any {
		return new(LauncherDo)
	},
}

// GetLauncherDo() 从对象池中获取LauncherDo
func GetLauncherDo() *LauncherDo {
	return poolLauncherDo.Get().(*LauncherDo)
}

// ReleaseLauncherDo 释放LauncherDo
func ReleaseLauncherDo(v *LauncherDo) {
	v.Items = v.Items[:0]
	v.Property = nil
	v.Version = nil
	poolLauncherDo.Put(v)
}
