package legalsuit

import (
	"sync"
)

// Attachmentlist 结构体
type Attachmentlist struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 文件key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
}

var poolAttachmentlist = sync.Pool{
	New: func() any {
		return new(Attachmentlist)
	},
}

// GetAttachmentlist() 从对象池中获取Attachmentlist
func GetAttachmentlist() *Attachmentlist {
	return poolAttachmentlist.Get().(*Attachmentlist)
}

// ReleaseAttachmentlist 释放Attachmentlist
func ReleaseAttachmentlist(v *Attachmentlist) {
	v.Name = ""
	v.Key = ""
	poolAttachmentlist.Put(v)
}
