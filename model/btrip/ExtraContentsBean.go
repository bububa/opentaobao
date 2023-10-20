package btrip

import (
	"sync"
)

// ExtraContentsBean 结构体
type ExtraContentsBean struct {
	// 说明内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 内容标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
}

var poolExtraContentsBean = sync.Pool{
	New: func() any {
		return new(ExtraContentsBean)
	},
}

// GetExtraContentsBean() 从对象池中获取ExtraContentsBean
func GetExtraContentsBean() *ExtraContentsBean {
	return poolExtraContentsBean.Get().(*ExtraContentsBean)
}

// ReleaseExtraContentsBean 释放ExtraContentsBean
func ReleaseExtraContentsBean(v *ExtraContentsBean) {
	v.Content = ""
	v.Title = ""
	poolExtraContentsBean.Put(v)
}
