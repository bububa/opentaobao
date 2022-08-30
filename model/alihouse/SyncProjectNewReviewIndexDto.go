package alihouse

// SyncProjectNewReviewIndexDto 结构体
type SyncProjectNewReviewIndexDto struct {
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 楼盘区位分数
	LocationScore string `json:"location_score,omitempty" xml:"location_score,omitempty"`
	// 楼盘交通分数
	TrafficScore string `json:"traffic_score,omitempty" xml:"traffic_score,omitempty"`
	// 楼盘生产力分数
	ProductForceScore string `json:"product_force_score,omitempty" xml:"product_force_score,omitempty"`
	// 楼盘配套分数
	MatchedScore string `json:"matched_score,omitempty" xml:"matched_score,omitempty"`
	// 楼盘户型分数
	HouseTypeScore string `json:"house_type_score,omitempty" xml:"house_type_score,omitempty"`
	// 楼盘总分
	HouseScore string `json:"house_score,omitempty" xml:"house_score,omitempty"`
	// 楼盘区位星级
	LocationLevel string `json:"location_level,omitempty" xml:"location_level,omitempty"`
	// 楼盘交通星级
	TrafficLevel string `json:"traffic_level,omitempty" xml:"traffic_level,omitempty"`
	// 楼盘产品力星级
	ProductForceLevel string `json:"product_force_level,omitempty" xml:"product_force_level,omitempty"`
	// 楼盘配套星级
	MatchedLevel string `json:"matched_level,omitempty" xml:"matched_level,omitempty"`
	// 楼盘户型星级
	HouseTypeLevel string `json:"house_type_level,omitempty" xml:"house_type_level,omitempty"`
	// 楼盘星级
	HouseLevel string `json:"house_level,omitempty" xml:"house_level,omitempty"`
	// 外部指数id
	OuterIndexId int64 `json:"outer_index_id,omitempty" xml:"outer_index_id,omitempty"`
	// 城市id
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 1有效0无效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
