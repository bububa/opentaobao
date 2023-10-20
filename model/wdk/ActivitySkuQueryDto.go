package wdk

import (
	"sync"
)

// ActivitySkuQueryDto 结构体
type ActivitySkuQueryDto struct {
	// 当前页码，从1开始
	Current int64 `json:"current,omitempty" xml:"current,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolActivitySkuQueryDto = sync.Pool{
	New: func() any {
		return new(ActivitySkuQueryDto)
	},
}

// GetActivitySkuQueryDto() 从对象池中获取ActivitySkuQueryDto
func GetActivitySkuQueryDto() *ActivitySkuQueryDto {
	return poolActivitySkuQueryDto.Get().(*ActivitySkuQueryDto)
}

// ReleaseActivitySkuQueryDto 释放ActivitySkuQueryDto
func ReleaseActivitySkuQueryDto(v *ActivitySkuQueryDto) {
	v.Current = 0
	v.PageSize = 0
	poolActivitySkuQueryDto.Put(v)
}
