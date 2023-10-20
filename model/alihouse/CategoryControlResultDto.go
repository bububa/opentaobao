package alihouse

import (
	"sync"
)

// CategoryControlResultDto 结构体
type CategoryControlResultDto struct {
	// 原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
}

var poolCategoryControlResultDto = sync.Pool{
	New: func() any {
		return new(CategoryControlResultDto)
	},
}

// GetCategoryControlResultDto() 从对象池中获取CategoryControlResultDto
func GetCategoryControlResultDto() *CategoryControlResultDto {
	return poolCategoryControlResultDto.Get().(*CategoryControlResultDto)
}

// ReleaseCategoryControlResultDto 释放CategoryControlResultDto
func ReleaseCategoryControlResultDto(v *CategoryControlResultDto) {
	v.Reason = ""
	v.CategoryId = 0
	poolCategoryControlResultDto.Put(v)
}
