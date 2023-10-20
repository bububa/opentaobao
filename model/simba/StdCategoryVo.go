package simba

import (
	"sync"
)

// StdCategoryVo 结构体
type StdCategoryVo struct {
	// 类目名称
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// 类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 类目级别，一级类目、二级目录
	CatLevel int64 `json:"cat_level,omitempty" xml:"cat_level,omitempty"`
}

var poolStdCategoryVo = sync.Pool{
	New: func() any {
		return new(StdCategoryVo)
	},
}

// GetStdCategoryVo() 从对象池中获取StdCategoryVo
func GetStdCategoryVo() *StdCategoryVo {
	return poolStdCategoryVo.Get().(*StdCategoryVo)
}

// ReleaseStdCategoryVo 释放StdCategoryVo
func ReleaseStdCategoryVo(v *StdCategoryVo) {
	v.CatName = ""
	v.CatId = 0
	v.CatLevel = 0
	poolStdCategoryVo.Put(v)
}
