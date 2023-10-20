package tmallitem

import (
	"sync"
)

// TmallBrand 结构体
type TmallBrand struct {
	// 搜索品牌名字
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 搜索品牌id
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
}

var poolTmallBrand = sync.Pool{
	New: func() any {
		return new(TmallBrand)
	},
}

// GetTmallBrand() 从对象池中获取TmallBrand
func GetTmallBrand() *TmallBrand {
	return poolTmallBrand.Get().(*TmallBrand)
}

// ReleaseTmallBrand 释放TmallBrand
func ReleaseTmallBrand(v *TmallBrand) {
	v.BrandName = ""
	v.BrandId = 0
	poolTmallBrand.Put(v)
}
