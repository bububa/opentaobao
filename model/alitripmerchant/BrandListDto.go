package alitripmerchant

import (
	"sync"
)

// BrandListDto 结构体
type BrandListDto struct {
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 品牌类型
	Species string `json:"species,omitempty" xml:"species,omitempty"`
	// 品牌地址
	Logo string `json:"logo,omitempty" xml:"logo,omitempty"`
	// 品牌编码
	BrandCode string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	// 子品牌编码
	SubType int64 `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
}

var poolBrandListDto = sync.Pool{
	New: func() any {
		return new(BrandListDto)
	},
}

// GetBrandListDto() 从对象池中获取BrandListDto
func GetBrandListDto() *BrandListDto {
	return poolBrandListDto.Get().(*BrandListDto)
}

// ReleaseBrandListDto 释放BrandListDto
func ReleaseBrandListDto(v *BrandListDto) {
	v.BrandName = ""
	v.Species = ""
	v.Logo = ""
	v.BrandCode = ""
	v.SubType = 0
	poolBrandListDto.Put(v)
}
