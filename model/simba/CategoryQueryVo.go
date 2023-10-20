package simba

import (
	"sync"
)

// CategoryQueryVo 结构体
type CategoryQueryVo struct {
	// 类目名称
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// 类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 类目级别，一级类目、二级目录
	CatLevel int64 `json:"cat_level,omitempty" xml:"cat_level,omitempty"`
}

var poolCategoryQueryVo = sync.Pool{
	New: func() any {
		return new(CategoryQueryVo)
	},
}

// GetCategoryQueryVo() 从对象池中获取CategoryQueryVo
func GetCategoryQueryVo() *CategoryQueryVo {
	return poolCategoryQueryVo.Get().(*CategoryQueryVo)
}

// ReleaseCategoryQueryVo 释放CategoryQueryVo
func ReleaseCategoryQueryVo(v *CategoryQueryVo) {
	v.CatName = ""
	v.CatId = 0
	v.CatLevel = 0
	poolCategoryQueryVo.Put(v)
}
