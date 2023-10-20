package alihouse

// UpdateNewHomeEcodeInfoDto 结构体
type UpdateNewHomeEcodeInfoDto struct {
	// 外部小区id
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 房源E码
	Ecode string `json:"ecode,omitempty" xml:"ecode,omitempty"`
	// 外部房源id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部项目店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 业务类型，1-新房，2-二手房，3-租房，默认为2
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 0-正常房源 1-临时房源（交易专用）
	HouseType int64 `json:"house_type,omitempty" xml:"house_type,omitempty"`
	// 是否为货 0-非货,信息流 1-货�商品
	IsCargo int64 `json:"is_cargo,omitempty" xml:"is_cargo,omitempty"`
}
