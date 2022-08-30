package alihouse

// SyncExistingHouseThreeDimensionalDto 结构体
type SyncExistingHouseThreeDimensionalDto struct {
	// 外部小区id
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 外部房源id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 3D户型编码
	UniqueNo string `json:"unique_no,omitempty" xml:"unique_no,omitempty"`
	// 封面图
	CoverPictureUrl string `json:"cover_picture_url,omitempty" xml:"cover_picture_url,omitempty"`
	// 3D数据包
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}
