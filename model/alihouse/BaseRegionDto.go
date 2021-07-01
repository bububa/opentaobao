package alihouse

// BaseRegionDto 结构体
type BaseRegionDto struct {
	// 高德围栏
	Fencing string `json:"fencing,omitempty" xml:"fencing,omitempty"`
	// 名称全称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 高德中心纬度
	GaodeLatitude string `json:"gaode_latitude,omitempty" xml:"gaode_latitude,omitempty"`
	// 高德中心经度
	GaodeLongitude string `json:"gaode_longitude,omitempty" xml:"gaode_longitude,omitempty"`
	// 名称简称
	NameSimple string `json:"name_simple,omitempty" xml:"name_simple,omitempty"`
	// 名称拼音全字母
	NamePinyin string `json:"name_pinyin,omitempty" xml:"name_pinyin,omitempty"`
	// 名称拼音首字母
	NamePinyinSimple string `json:"name_pinyin_simple,omitempty" xml:"name_pinyin_simple,omitempty"`
	// 区域类型
	RegionType string `json:"region_type,omitempty" xml:"region_type,omitempty"`
	// 区域代码
	RegionId int64 `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// 外部城区ID -- 唯一
	OuterRegionId string `json:"outer_region_id,omitempty" xml:"outer_region_id,omitempty"`
	// 1或0 1删除 0不删
	IsDeleted string `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 父级ID
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
}
