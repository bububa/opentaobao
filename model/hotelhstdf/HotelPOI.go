package hotelhstdf

// HotelPOI 结构体
type HotelPOI struct {
	// 暂不使用
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// poi location纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// poi locaiton经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 中文名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 所在国家英文名
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 行政区划树，冗余字段，不一定有值
	DivisionTree string `json:"division_tree,omitempty" xml:"division_tree,omitempty"`
	// 英文名
	EnName string `json:"en_name,omitempty" xml:"en_name,omitempty"`
	// 时区，冗余字段，不一定有值
	TimeZoneId string `json:"time_zone_id,omitempty" xml:"time_zone_id,omitempty"`
	// 搜索热度
	SearchPercent string `json:"search_percent,omitempty" xml:"search_percent,omitempty"`
	// 所在城市id
	City int64 `json:"city,omitempty" xml:"city,omitempty"`
	// 热度
	Hot int64 `json:"hot,omitempty" xml:"hot,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 分组内排序
	OrderIndex int64 `json:"order_index,omitempty" xml:"order_index,omitempty"`
	// 父id
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 类型，3--地铁线路,4--地铁站,20--考点,5--景点,6--机场,7--火车站,8--汽车站,9--医院,10--大学,11--热点搜索,12--城市,13--办公区
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 所在国家id
	CountryCode int64 `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 所在区县id
	District int64 `json:"district,omitempty" xml:"district,omitempty"`
	// 能精确到的最低层级区划id，优先区县，次选城市
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
	// 国内外，0-国内，1--国外
	Domestic int64 `json:"domestic,omitempty" xml:"domestic,omitempty"`
	// 经纬度服务商，0--高德，一般用于国内，1- - 谷歌，一般用于国际
	PositionType int64 `json:"position_type,omitempty" xml:"position_type,omitempty"`
	// 所在省份id，北京直辖市
	Province int64 `json:"province,omitempty" xml:"province,omitempty"`
	// 最精确的行政区划对应的平台id
	TrdiDivisionId int64 `json:"trdi_division_id,omitempty" xml:"trdi_division_id,omitempty"`
}
