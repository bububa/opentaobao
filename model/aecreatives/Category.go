package aecreatives

import (
	"sync"
)

// Category 结构体
type Category struct {
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 类目ID
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 父类目ID
	ParentCategoryId int64 `json:"parent_category_id,omitempty" xml:"parent_category_id,omitempty"`
}

var poolCategory = sync.Pool{
	New: func() any {
		return new(Category)
	},
}

// GetCategory() 从对象池中获取Category
func GetCategory() *Category {
	return poolCategory.Get().(*Category)
}

// ReleaseCategory 释放Category
func ReleaseCategory(v *Category) {
	v.CategoryName = ""
	v.CategoryId = 0
	v.ParentCategoryId = 0
	poolCategory.Put(v)
}
