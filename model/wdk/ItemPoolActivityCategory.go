package wdk

import (
	"sync"
)

// ItemPoolActivityCategory 结构体
type ItemPoolActivityCategory struct {
	// 类目分组
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 分组id
	LogicGroupNumber int64 `json:"logic_group_number,omitempty" xml:"logic_group_number,omitempty"`
}

var poolItemPoolActivityCategory = sync.Pool{
	New: func() any {
		return new(ItemPoolActivityCategory)
	},
}

// GetItemPoolActivityCategory() 从对象池中获取ItemPoolActivityCategory
func GetItemPoolActivityCategory() *ItemPoolActivityCategory {
	return poolItemPoolActivityCategory.Get().(*ItemPoolActivityCategory)
}

// ReleaseItemPoolActivityCategory 释放ItemPoolActivityCategory
func ReleaseItemPoolActivityCategory(v *ItemPoolActivityCategory) {
	v.CategoryId = ""
	v.LogicGroupNumber = 0
	poolItemPoolActivityCategory.Put(v)
}
