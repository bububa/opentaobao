package alihouse

import (
	"sync"
)

// BaseRegionDto 结构体
type BaseRegionDto struct {
	// 高德围栏
	Fencing string `json:"fencing,omitempty" xml:"fencing,omitempty"`
	// 高德中心纬度
	GaodeLatitude string `json:"gaode_latitude,omitempty" xml:"gaode_latitude,omitempty"`
	// 高德中心经度
	GaodeLongitude string `json:"gaode_longitude,omitempty" xml:"gaode_longitude,omitempty"`
	// 名称简称
	NameSimple string `json:"name_simple,omitempty" xml:"name_simple,omitempty"`
	// 名称全称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 名称拼音全字母
	NamePinyin string `json:"name_pinyin,omitempty" xml:"name_pinyin,omitempty"`
	// 名称拼音首字母
	NamePinyinSimple string `json:"name_pinyin_simple,omitempty" xml:"name_pinyin_simple,omitempty"`
	// 区域类型
	RegionType string `json:"region_type,omitempty" xml:"region_type,omitempty"`
	// 外部城区ID -- 唯一
	OuterRegionId string `json:"outer_region_id,omitempty" xml:"outer_region_id,omitempty"`
	// 1或0 1删除 0不删
	IsDeleted string `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 区域代码
	RegionId int64 `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// 父级ID
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 数据源类型（1-新房 2-二手房）
	SourceType int64 `json:"source_type,omitempty" xml:"source_type,omitempty"`
}

var poolBaseRegionDto = sync.Pool{
	New: func() any {
		return new(BaseRegionDto)
	},
}

// GetBaseRegionDto() 从对象池中获取BaseRegionDto
func GetBaseRegionDto() *BaseRegionDto {
	return poolBaseRegionDto.Get().(*BaseRegionDto)
}

// ReleaseBaseRegionDto 释放BaseRegionDto
func ReleaseBaseRegionDto(v *BaseRegionDto) {
	v.Fencing = ""
	v.GaodeLatitude = ""
	v.GaodeLongitude = ""
	v.NameSimple = ""
	v.Name = ""
	v.NamePinyin = ""
	v.NamePinyinSimple = ""
	v.RegionType = ""
	v.OuterRegionId = ""
	v.IsDeleted = ""
	v.RegionId = 0
	v.ParentId = 0
	v.SourceType = 0
	poolBaseRegionDto.Put(v)
}
