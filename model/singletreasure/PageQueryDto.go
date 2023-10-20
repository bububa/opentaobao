package singletreasure

import (
	"sync"
)

// PageQueryDto 结构体
type PageQueryDto struct {
	// 查询条件请求体
	Query *ActivityInfoListQueryDto `json:"query,omitempty" xml:"query,omitempty"`
	// 页码,最大为50
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 页数，从1开始
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
}

var poolPageQueryDto = sync.Pool{
	New: func() any {
		return new(PageQueryDto)
	},
}

// GetPageQueryDto() 从对象池中获取PageQueryDto
func GetPageQueryDto() *PageQueryDto {
	return poolPageQueryDto.Get().(*PageQueryDto)
}

// ReleasePageQueryDto 释放PageQueryDto
func ReleasePageQueryDto(v *PageQueryDto) {
	v.Query = nil
	v.PageSize = 0
	v.PageNumber = 0
	poolPageQueryDto.Put(v)
}
