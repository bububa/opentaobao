package icburfq

import (
	"sync"
)

// Attachedfiles 结构体
type Attachedfiles struct {
	// 文件名
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 文件地址
	FileUrl string `json:"file_url,omitempty" xml:"file_url,omitempty"`
}

var poolAttachedfiles = sync.Pool{
	New: func() any {
		return new(Attachedfiles)
	},
}

// GetAttachedfiles() 从对象池中获取Attachedfiles
func GetAttachedfiles() *Attachedfiles {
	return poolAttachedfiles.Get().(*Attachedfiles)
}

// ReleaseAttachedfiles 释放Attachedfiles
func ReleaseAttachedfiles(v *Attachedfiles) {
	v.FileName = ""
	v.FileUrl = ""
	poolAttachedfiles.Put(v)
}
