package alihealthcrm

import (
	"sync"
)

// Contents 结构体
type Contents struct {
	// 商品列表
	Items []string `json:"items,omitempty" xml:"items>string,omitempty"`
	// 文章标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 文章标签
	Tags string `json:"tags,omitempty" xml:"tags,omitempty"`
	// 图片
	PhotoUrl string `json:"photo_url,omitempty" xml:"photo_url,omitempty"`
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 链接
	LinkUrl string `json:"link_url,omitempty" xml:"link_url,omitempty"`
	// publishTime
	PublishTime string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 商品计数
	ItemCount int64 `json:"item_count,omitempty" xml:"item_count,omitempty"`
}

var poolContents = sync.Pool{
	New: func() any {
		return new(Contents)
	},
}

// GetContents() 从对象池中获取Contents
func GetContents() *Contents {
	return poolContents.Get().(*Contents)
}

// ReleaseContents 释放Contents
func ReleaseContents(v *Contents) {
	v.Items = v.Items[:0]
	v.Title = ""
	v.Tags = ""
	v.PhotoUrl = ""
	v.Content = ""
	v.LinkUrl = ""
	v.PublishTime = ""
	v.ItemCount = 0
	poolContents.Put(v)
}
