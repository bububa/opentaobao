package icbuseller

import (
	"sync"
)

// PageDto 结构体
type PageDto struct {
	// 总数据量
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 每页显示数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}

var poolPageDto = sync.Pool{
	New: func() any {
		return new(PageDto)
	},
}

// GetPageDto() 从对象池中获取PageDto
func GetPageDto() *PageDto {
	return poolPageDto.Get().(*PageDto)
}

// ReleasePageDto 释放PageDto
func ReleasePageDto(v *PageDto) {
	v.TotalItem = 0
	v.TotalPage = 0
	v.PageSize = 0
	v.CurrentPage = 0
	poolPageDto.Put(v)
}
