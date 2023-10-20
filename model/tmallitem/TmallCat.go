package tmallitem

import (
	"sync"
)

// TmallCat 结构体
type TmallCat struct {
	// 搜索前台类目名字
	CatName string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// 搜索前台类目id
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
}

var poolTmallCat = sync.Pool{
	New: func() any {
		return new(TmallCat)
	},
}

// GetTmallCat() 从对象池中获取TmallCat
func GetTmallCat() *TmallCat {
	return poolTmallCat.Get().(*TmallCat)
}

// ReleaseTmallCat 释放TmallCat
func ReleaseTmallCat(v *TmallCat) {
	v.CatName = ""
	v.CatId = 0
	poolTmallCat.Put(v)
}
