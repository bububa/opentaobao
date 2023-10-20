package wdk

import (
	"sync"
)

// PageInfoDto 结构体
type PageInfoDto struct {
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 总页数
	Pages int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页数
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}

var poolPageInfoDto = sync.Pool{
	New: func() any {
		return new(PageInfoDto)
	},
}

// GetPageInfoDto() 从对象池中获取PageInfoDto
func GetPageInfoDto() *PageInfoDto {
	return poolPageInfoDto.Get().(*PageInfoDto)
}

// ReleasePageInfoDto 释放PageInfoDto
func ReleasePageInfoDto(v *PageInfoDto) {
	v.Total = 0
	v.Pages = 0
	v.PageSize = 0
	v.PageNum = 0
	poolPageInfoDto.Put(v)
}
