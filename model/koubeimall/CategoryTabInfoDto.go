package koubeimall

import (
	"sync"
)

// CategoryTabInfoDto 结构体
type CategoryTabInfoDto struct {
	// 前台类目ids
	CategoryIdList []string `json:"category_id_list,omitempty" xml:"category_id_list>string,omitempty"`
	// 前台类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
}

var poolCategoryTabInfoDto = sync.Pool{
	New: func() any {
		return new(CategoryTabInfoDto)
	},
}

// GetCategoryTabInfoDto() 从对象池中获取CategoryTabInfoDto
func GetCategoryTabInfoDto() *CategoryTabInfoDto {
	return poolCategoryTabInfoDto.Get().(*CategoryTabInfoDto)
}

// ReleaseCategoryTabInfoDto 释放CategoryTabInfoDto
func ReleaseCategoryTabInfoDto(v *CategoryTabInfoDto) {
	v.CategoryIdList = v.CategoryIdList[:0]
	v.CategoryName = ""
	poolCategoryTabInfoDto.Put(v)
}
