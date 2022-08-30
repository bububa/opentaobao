package alihouse

// HouseFeaturesDto 结构体
type HouseFeaturesDto struct {
	// 外部房源id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部小区id
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 租房业务模式
	HouseModel int64 `json:"house_model,omitempty" xml:"house_model,omitempty"`
}
