package ott

import (
	"sync"
)

// VersionDo 结构体
type VersionDo struct {
	// 桌面标识
	LauncherCode string `json:"launcher_code,omitempty" xml:"launcher_code,omitempty"`
	// 桌面名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 图标规格
	EntrySize int64 `json:"entry_size,omitempty" xml:"entry_size,omitempty"`
	// 桌面ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolVersionDo = sync.Pool{
	New: func() any {
		return new(VersionDo)
	},
}

// GetVersionDo() 从对象池中获取VersionDo
func GetVersionDo() *VersionDo {
	return poolVersionDo.Get().(*VersionDo)
}

// ReleaseVersionDo 释放VersionDo
func ReleaseVersionDo(v *VersionDo) {
	v.LauncherCode = ""
	v.Name = ""
	v.EntrySize = 0
	v.Id = 0
	poolVersionDo.Put(v)
}
