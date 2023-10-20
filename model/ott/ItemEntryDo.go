package ott

import (
	"sync"
)

// ItemEntryDo 结构体
type ItemEntryDo struct {
	// 行为扩展
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 跳转行为
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 入口图标
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 入口名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
	// 入口ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolItemEntryDo = sync.Pool{
	New: func() any {
		return new(ItemEntryDo)
	},
}

// GetItemEntryDo() 从对象池中获取ItemEntryDo
func GetItemEntryDo() *ItemEntryDo {
	return poolItemEntryDo.Get().(*ItemEntryDo)
}

// ReleaseItemEntryDo 释放ItemEntryDo
func ReleaseItemEntryDo(v *ItemEntryDo) {
	v.Extra = ""
	v.Action = ""
	v.PicUrl = ""
	v.Name = ""
	v.Sort = 0
	v.Id = 0
	poolItemEntryDo.Put(v)
}
