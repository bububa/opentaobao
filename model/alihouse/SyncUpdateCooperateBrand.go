package alihouse

// SyncUpdateCooperateBrand 结构体
type SyncUpdateCooperateBrand struct {
	// 外部小区id
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 外部房源id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部品牌id
	OuterBrandId string `json:"outer_brand_id,omitempty" xml:"outer_brand_id,omitempty"`
}
