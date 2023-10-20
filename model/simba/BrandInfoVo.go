package simba

import (
	"sync"
)

// BrandInfoVo 结构体
type BrandInfoVo struct {
	// 品牌ID
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
}

var poolBrandInfoVo = sync.Pool{
	New: func() any {
		return new(BrandInfoVo)
	},
}

// GetBrandInfoVo() 从对象池中获取BrandInfoVo
func GetBrandInfoVo() *BrandInfoVo {
	return poolBrandInfoVo.Get().(*BrandInfoVo)
}

// ReleaseBrandInfoVo 释放BrandInfoVo
func ReleaseBrandInfoVo(v *BrandInfoVo) {
	v.BrandId = ""
	v.BrandName = ""
	poolBrandInfoVo.Put(v)
}
