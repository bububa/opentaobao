package ott

import (
	"sync"
)

// Entrylist 结构体
type Entrylist struct {
	// 入口名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 入口图标
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 跳转行为
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 行为扩展
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 入口ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
}

var poolEntrylist = sync.Pool{
	New: func() any {
		return new(Entrylist)
	},
}

// GetEntrylist() 从对象池中获取Entrylist
func GetEntrylist() *Entrylist {
	return poolEntrylist.Get().(*Entrylist)
}

// ReleaseEntrylist 释放Entrylist
func ReleaseEntrylist(v *Entrylist) {
	v.Name = ""
	v.PicUrl = ""
	v.Action = ""
	v.Extra = ""
	v.Id = 0
	v.Sort = 0
	poolEntrylist.Put(v)
}
