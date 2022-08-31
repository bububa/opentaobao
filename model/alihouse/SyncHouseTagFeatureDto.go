package alihouse

// SyncHouseTagFeatureDto 结构体
type SyncHouseTagFeatureDto struct {
	// 实体类
	FeatureDtoList []FeatureDto `json:"feature_dto_list,omitempty" xml:"feature_dto_list>feature_dto,omitempty"`
	// 123
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 123
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 1
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 1
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 1
	OnlineStatus int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
}
