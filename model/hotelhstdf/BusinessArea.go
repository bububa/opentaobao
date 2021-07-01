package hotelhstdf

// BusinessArea 结构体
type BusinessArea struct {
	// 商圈围栏数据，经纬度坐标使用英文逗号拼接
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 商圈图片
	AreaPic string `json:"area_pic,omitempty" xml:"area_pic,omitempty"`
	// 所在城市id，北京市
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 所在国家英文名
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 所在区县id
	District int64 `json:"district,omitempty" xml:"district,omitempty"`
	// 国内外，0-国内，1--国外
	Domestic int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 计算酒店所属商圈是否需要外延，0-不需要，1-需要
	ExtentMeter int64 `json:"extent_meter,omitempty" xml:"extent_meter,omitempty"`
	// 热度
	Hot int64 `json:"hot,omitempty" xml:"hot,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 商圈中心点纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 商圈中心点经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 暂不使用
	Major int64 `json:"major,omitempty" xml:"major,omitempty"`
	// 中文名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 暂不使用
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 所在省份id，北京直辖市
	Province int64 `json:"province,omitempty" xml:"province,omitempty"`
	// 商圈跨多个地区时，其他地区id适应英文逗号拼接
	Region string `json:"region,omitempty" xml:"region,omitempty"`
	// 暂不使用
	Source int64 `json:"source,omitempty" xml:"source,omitempty"`
	// 商圈推荐导语
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 商圈标签，英文逗号分隔，注意与酒店标签不是同一类概念
	Tags []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
	// 商圈缩略图
	AreaPicNail string `json:"area_pic_nail,omitempty" xml:"area_pic_nail,omitempty"`
	// 国家编码，取代country使用
	CountryCode int64 `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 能精确到的最低层级区划id，优先区县，次选城市
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
	// 行政区划树，冗余字段，不一定有值
	DivisionTree string `json:"division_tree,omitempty" xml:"division_tree,omitempty"`
	// 商圈英文名
	EnName string `json:"en_name,omitempty" xml:"en_name,omitempty"`
	// 经纬度服务商，0--高德，一般用于国内，1-- 谷歌，一般用于国际
	PositionType int64 `json:"position_type,omitempty" xml:"position_type,omitempty"`
	// 时区，冗余字段，不一定有值
	TimeZoneId string `json:"time_zone_id,omitempty" xml:"time_zone_id,omitempty"`
	// 最精确的行政区划对应的平台id
	TrdiDivisionId int64 `json:"trdi_division_id,omitempty" xml:"trdi_division_id,omitempty"`
}
