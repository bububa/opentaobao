package icburfq

import (
	"sync"
)

// Categorylist 结构体
type Categorylist struct {
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 类目ID
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolCategorylist = sync.Pool{
	New: func() any {
		return new(Categorylist)
	},
}

// GetCategorylist() 从对象池中获取Categorylist
func GetCategorylist() *Categorylist {
	return poolCategorylist.Get().(*Categorylist)
}

// ReleaseCategorylist 释放Categorylist
func ReleaseCategorylist(v *Categorylist) {
	v.CategoryName = ""
	v.CategoryId = 0
	v.Count = 0
	poolCategorylist.Put(v)
}
