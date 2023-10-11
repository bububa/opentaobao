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
	// 跳转链接
	JumpUrl string `json:"jump_url,omitempty" xml:"jump_url,omitempty"`
	// 类型 1：躺平3D户型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
