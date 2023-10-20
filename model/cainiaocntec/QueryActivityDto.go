package cainiaocntec

import (
	"sync"
)

// QueryActivityDto 结构体
type QueryActivityDto struct {
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 是否分页
	Page bool `json:"page,omitempty" xml:"page,omitempty"`
}

var poolQueryActivityDto = sync.Pool{
	New: func() any {
		return new(QueryActivityDto)
	},
}

// GetQueryActivityDto() 从对象池中获取QueryActivityDto
func GetQueryActivityDto() *QueryActivityDto {
	return poolQueryActivityDto.Get().(*QueryActivityDto)
}

// ReleaseQueryActivityDto 释放QueryActivityDto
func ReleaseQueryActivityDto(v *QueryActivityDto) {
	v.StoreId = ""
	v.PageSize = 0
	v.PageIndex = 0
	v.Page = false
	poolQueryActivityDto.Put(v)
}
