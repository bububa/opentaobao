package viapi

import (
	"sync"
)

// Category 结构体
type Category struct {
	// 分类ID的匹配度，大于0小于等于1
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 类目ID
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
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
	v.Score = ""
	v.CategoryName = ""
	v.CategoryId = ""
	poolCategory.Put(v)
}
