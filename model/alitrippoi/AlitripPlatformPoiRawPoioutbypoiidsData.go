package alitrippoi

// AlitripPlatformPoiRawPoioutbypoiidsData 结构体
type AlitripPlatformPoiRawPoioutbypoiidsData struct {
	// poiId
	SourceId string `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 行政区划树名
	DivisionTreeName string `json:"division_tree_name,omitempty" xml:"division_tree_name,omitempty"`
	// 行政区划树
	DivisionTreeId string `json:"division_tree_id,omitempty" xml:"division_tree_id,omitempty"`
	// 开放时间
	OpenTime string `json:"open_time,omitempty" xml:"open_time,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 经度
	Lon string `json:"lon,omitempty" xml:"lon,omitempty"`
	// 电话
	Telephone string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 类目
	FirstCategory int64 `json:"first_category,omitempty" xml:"first_category,omitempty"`
	// poi英文名
	NameEn string `json:"name_en,omitempty" xml:"name_en,omitempty"`
	// poi名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// poiId
	PoiId int64 `json:"poi_id,omitempty" xml:"poi_id,omitempty"`
}
