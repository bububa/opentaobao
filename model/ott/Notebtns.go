package ott

import (
	"sync"
)

// Notebtns 结构体
type Notebtns struct {
	// 便签类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 便签名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
}

var poolNotebtns = sync.Pool{
	New: func() any {
		return new(Notebtns)
	},
}

// GetNotebtns() 从对象池中获取Notebtns
func GetNotebtns() *Notebtns {
	return poolNotebtns.Get().(*Notebtns)
}

// ReleaseNotebtns 释放Notebtns
func ReleaseNotebtns(v *Notebtns) {
	v.Type = ""
	v.Name = ""
	v.Sort = 0
	poolNotebtns.Put(v)
}
