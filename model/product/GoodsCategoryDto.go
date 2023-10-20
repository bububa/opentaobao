package product

import (
	"sync"
)

// GoodsCategoryDto 结构体
type GoodsCategoryDto struct {
	// 二级类目名称
	SecondCategoryName string `json:"second_category_name,omitempty" xml:"second_category_name,omitempty"`
	// 一级类目名称
	FirstCategoryName string `json:"first_category_name,omitempty" xml:"first_category_name,omitempty"`
	// 二级类目ID
	SecondCategoryId int64 `json:"second_category_id,omitempty" xml:"second_category_id,omitempty"`
	// 一级类目ID
	FirstCategoryId int64 `json:"first_category_id,omitempty" xml:"first_category_id,omitempty"`
}

var poolGoodsCategoryDto = sync.Pool{
	New: func() any {
		return new(GoodsCategoryDto)
	},
}

// GetGoodsCategoryDto() 从对象池中获取GoodsCategoryDto
func GetGoodsCategoryDto() *GoodsCategoryDto {
	return poolGoodsCategoryDto.Get().(*GoodsCategoryDto)
}

// ReleaseGoodsCategoryDto 释放GoodsCategoryDto
func ReleaseGoodsCategoryDto(v *GoodsCategoryDto) {
	v.SecondCategoryName = ""
	v.FirstCategoryName = ""
	v.SecondCategoryId = 0
	v.FirstCategoryId = 0
	poolGoodsCategoryDto.Put(v)
}
