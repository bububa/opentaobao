package alihouse

import (
	"sync"
)

// UpdateECodeBuildingDto 结构体
type UpdateECodeBuildingDto struct {
	// 外部私域楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部货-楼栋id（外部唯一识别码）
	OuterTid string `json:"outer_tid,omitempty" xml:"outer_tid,omitempty"`
	// 外部项目店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 楼栋E码
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
}

var poolUpdateECodeBuildingDto = sync.Pool{
	New: func() any {
		return new(UpdateECodeBuildingDto)
	},
}

// GetUpdateECodeBuildingDto() 从对象池中获取UpdateECodeBuildingDto
func GetUpdateECodeBuildingDto() *UpdateECodeBuildingDto {
	return poolUpdateECodeBuildingDto.Get().(*UpdateECodeBuildingDto)
}

// ReleaseUpdateECodeBuildingDto 释放UpdateECodeBuildingDto
func ReleaseUpdateECodeBuildingDto(v *UpdateECodeBuildingDto) {
	v.OuterId = ""
	v.OuterTid = ""
	v.OuterStoreId = ""
	v.ECode = ""
	poolUpdateECodeBuildingDto.Put(v)
}
