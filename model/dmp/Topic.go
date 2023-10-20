package dmp

import (
	"sync"
)

// Topic 结构体
type Topic struct {
	// 榜单下的模版对象数组
	Templates []Template `json:"templates,omitempty" xml:"templates>template,omitempty"`
	// 榜单名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 榜单描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 榜单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolTopic = sync.Pool{
	New: func() any {
		return new(Topic)
	},
}

// GetTopic() 从对象池中获取Topic
func GetTopic() *Topic {
	return poolTopic.Get().(*Topic)
}

// ReleaseTopic 释放Topic
func ReleaseTopic(v *Topic) {
	v.Templates = v.Templates[:0]
	v.Name = ""
	v.Description = ""
	v.Id = 0
	poolTopic.Put(v)
}
