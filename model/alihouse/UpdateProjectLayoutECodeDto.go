package alihouse

import (
	"sync"
)

// UpdateProjectLayoutECodeDto 结构体
type UpdateProjectLayoutECodeDto struct {
	// 楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部货ID
	OuterTid string `json:"outer_tid,omitempty" xml:"outer_tid,omitempty"`
	// 户型外部唯一ID
	OuterLayoutId string `json:"outer_layout_id,omitempty" xml:"outer_layout_id,omitempty"`
	// 外部项目店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// eCode
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 数据源类型 （1-新房 2-二手房）
	SourceType int64 `json:"source_type,omitempty" xml:"source_type,omitempty"`
}

var poolUpdateProjectLayoutECodeDto = sync.Pool{
	New: func() any {
		return new(UpdateProjectLayoutECodeDto)
	},
}

// GetUpdateProjectLayoutECodeDto() 从对象池中获取UpdateProjectLayoutECodeDto
func GetUpdateProjectLayoutECodeDto() *UpdateProjectLayoutECodeDto {
	return poolUpdateProjectLayoutECodeDto.Get().(*UpdateProjectLayoutECodeDto)
}

// ReleaseUpdateProjectLayoutECodeDto 释放UpdateProjectLayoutECodeDto
func ReleaseUpdateProjectLayoutECodeDto(v *UpdateProjectLayoutECodeDto) {
	v.OuterId = ""
	v.OuterTid = ""
	v.OuterLayoutId = ""
	v.OuterStoreId = ""
	v.ECode = ""
	v.SourceType = 0
	poolUpdateProjectLayoutECodeDto.Put(v)
}
