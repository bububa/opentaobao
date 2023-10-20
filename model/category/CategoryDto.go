package category

import (
	"sync"
)

// CategoryDto 结构体
type CategoryDto struct {
	// 子节点
	Childrens []CategoryDto `json:"childrens,omitempty" xml:"childrens>category_dto,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 类目ID
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
}

var poolCategoryDto = sync.Pool{
	New: func() any {
		return new(CategoryDto)
	},
}

// GetCategoryDto() 从对象池中获取CategoryDto
func GetCategoryDto() *CategoryDto {
	return poolCategoryDto.Get().(*CategoryDto)
}

// ReleaseCategoryDto 释放CategoryDto
func ReleaseCategoryDto(v *CategoryDto) {
	v.Childrens = v.Childrens[:0]
	v.CategoryName = ""
	v.CategoryId = ""
	poolCategoryDto.Put(v)
}
