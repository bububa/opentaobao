package aesolution

import (
	"sync"
)

// CategoryInfo 结构体
type CategoryInfo struct {
	// multi langauge names of the categories
	MultiLanguageNames string `json:"multi_language_names,omitempty" xml:"multi_language_names,omitempty"`
	// category id
	ChildrenCategoryId int64 `json:"children_category_id,omitempty" xml:"children_category_id,omitempty"`
	// level of the categories. As for root categories, the level is 1
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
	// whether the category is leaf or not
	IsLeafCategory bool `json:"is_leaf_category,omitempty" xml:"is_leaf_category,omitempty"`
}

var poolCategoryInfo = sync.Pool{
	New: func() any {
		return new(CategoryInfo)
	},
}

// GetCategoryInfo() 从对象池中获取CategoryInfo
func GetCategoryInfo() *CategoryInfo {
	return poolCategoryInfo.Get().(*CategoryInfo)
}

// ReleaseCategoryInfo 释放CategoryInfo
func ReleaseCategoryInfo(v *CategoryInfo) {
	v.MultiLanguageNames = ""
	v.ChildrenCategoryId = 0
	v.Level = 0
	v.IsLeafCategory = false
	poolCategoryInfo.Put(v)
}
