package aeusergrowth

import (
	"sync"
)

// Ext 结构体
type Ext struct {
	// search page link
	SearchPageLink string `json:"search_page_link,omitempty" xml:"search_page_link,omitempty"`
}

var poolExt = sync.Pool{
	New: func() any {
		return new(Ext)
	},
}

// GetExt() 从对象池中获取Ext
func GetExt() *Ext {
	return poolExt.Get().(*Ext)
}

// ReleaseExt 释放Ext
func ReleaseExt(v *Ext) {
	v.SearchPageLink = ""
	poolExt.Put(v)
}
