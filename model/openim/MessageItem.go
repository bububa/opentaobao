package openim

import (
	"sync"
)

// MessageItem 结构体
type MessageItem struct {
	// 节点类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 节点值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolMessageItem = sync.Pool{
	New: func() any {
		return new(MessageItem)
	},
}

// GetMessageItem() 从对象池中获取MessageItem
func GetMessageItem() *MessageItem {
	return poolMessageItem.Get().(*MessageItem)
}

// ReleaseMessageItem 释放MessageItem
func ReleaseMessageItem(v *MessageItem) {
	v.Type = ""
	v.Value = ""
	poolMessageItem.Put(v)
}
