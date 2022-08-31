package alihouse

// BaseBusinessDto 结构体
type BaseBusinessDto struct {
	// 高的围栏
	Fencing string `json:"fencing,omitempty" xml:"fencing,omitempty"`
	// 高德中心纬度坐标
	GaodeLatitude string `json:"gaode_latitude,omitempty" xml:"gaode_latitude,omitempty"`
	// 高德中心经度坐标
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
	// 外部ID -- 唯一
	OuterBusinessId string `json:"outer_business_id,omitempty" xml:"outer_business_id,omitempty"`
	// 是否删除 1 删除 0 不删（默认为0）
	IsDeleted string `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	// 区域代码
	RegionId int64 `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// 城市代码
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 数据源类型（1-新房 2-二手房）
	SourceType int64 `json:"source_type,omitempty" xml:"source_type,omitempty"`
}
