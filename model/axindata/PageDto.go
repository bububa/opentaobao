package axindata

import (
	"sync"
)

// PageDto 结构体
type PageDto struct {
	// 标准酒店id列表
	DataList []int64 `json:"data_list,omitempty" xml:"data_list>int64,omitempty"`
	// 记录数
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 页码
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 上次返回最大酒店id
	MaxDateId int64 `json:"max_date_id,omitempty" xml:"max_date_id,omitempty"`
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
	v.DataList = v.DataList[:0]
	v.Count = 0
	v.PageNo = 0
	v.PageSize = 0
	v.MaxDateId = 0
	poolPageDto.Put(v)
}
