package aesolution

import (
	"sync"
)

// SuggestCategory 结构体
type SuggestCategory struct {
	// category id path
	IdPath []int64 `json:"id_path,omitempty" xml:"id_path>int64,omitempty"`
	// category name path
	NamePath []string `json:"name_path,omitempty" xml:"name_path>string,omitempty"`
	// category name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// category id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolSuggestCategory = sync.Pool{
	New: func() any {
		return new(SuggestCategory)
	},
}

// GetSuggestCategory() 从对象池中获取SuggestCategory
func GetSuggestCategory() *SuggestCategory {
	return poolSuggestCategory.Get().(*SuggestCategory)
}

// ReleaseSuggestCategory 释放SuggestCategory
func ReleaseSuggestCategory(v *SuggestCategory) {
	v.IdPath = v.IdPath[:0]
	v.NamePath = v.NamePath[:0]
	v.Name = ""
	v.Id = 0
	poolSuggestCategory.Put(v)
}
