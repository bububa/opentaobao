package xhotelitem

import (
	"sync"
)

// TagEntityDoList 结构体
type TagEntityDoList struct {
	// 实体id
	EntityId int64 `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
}

var poolTagEntityDoList = sync.Pool{
	New: func() any {
		return new(TagEntityDoList)
	},
}

// GetTagEntityDoList() 从对象池中获取TagEntityDoList
func GetTagEntityDoList() *TagEntityDoList {
	return poolTagEntityDoList.Get().(*TagEntityDoList)
}

// ReleaseTagEntityDoList 释放TagEntityDoList
func ReleaseTagEntityDoList(v *TagEntityDoList) {
	v.EntityId = 0
	poolTagEntityDoList.Put(v)
}
