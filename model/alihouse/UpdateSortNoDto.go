package alihouse

import (
	"sync"
)

// UpdateSortNoDto 结构体
type UpdateSortNoDto struct {
	// 楼盘外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 排序值 大于等于10的数字
	SortNo int64 `json:"sort_no,omitempty" xml:"sort_no,omitempty"`
}

var poolUpdateSortNoDto = sync.Pool{
	New: func() any {
		return new(UpdateSortNoDto)
	},
}

// GetUpdateSortNoDto() 从对象池中获取UpdateSortNoDto
func GetUpdateSortNoDto() *UpdateSortNoDto {
	return poolUpdateSortNoDto.Get().(*UpdateSortNoDto)
}

// ReleaseUpdateSortNoDto 释放UpdateSortNoDto
func ReleaseUpdateSortNoDto(v *UpdateSortNoDto) {
	v.OuterId = ""
	v.SortNo = 0
	poolUpdateSortNoDto.Put(v)
}
