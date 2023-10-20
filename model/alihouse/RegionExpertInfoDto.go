package alihouse

import (
	"sync"
)

// RegionExpertInfoDto 结构体
type RegionExpertInfoDto struct {
	// 外部经纪人ID
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 经纪人状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolRegionExpertInfoDto = sync.Pool{
	New: func() any {
		return new(RegionExpertInfoDto)
	},
}

// GetRegionExpertInfoDto() 从对象池中获取RegionExpertInfoDto
func GetRegionExpertInfoDto() *RegionExpertInfoDto {
	return poolRegionExpertInfoDto.Get().(*RegionExpertInfoDto)
}

// ReleaseRegionExpertInfoDto 释放RegionExpertInfoDto
func ReleaseRegionExpertInfoDto(v *RegionExpertInfoDto) {
	v.OuterConsultantId = ""
	v.OuterStoreId = ""
	v.Status = 0
	poolRegionExpertInfoDto.Put(v)
}
