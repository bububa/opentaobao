package tvupadmin

import (
	"sync"
)

// BrandDo 结构体
type BrandDo struct {
	// brandName
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// brandId
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
}

var poolBrandDo = sync.Pool{
	New: func() any {
		return new(BrandDo)
	},
}

// GetBrandDo() 从对象池中获取BrandDo
func GetBrandDo() *BrandDo {
	return poolBrandDo.Get().(*BrandDo)
}

// ReleaseBrandDo 释放BrandDo
func ReleaseBrandDo(v *BrandDo) {
	v.BrandName = ""
	v.BrandId = 0
	poolBrandDo.Put(v)
}
