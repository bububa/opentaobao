package alitripmerchant

// RoomDetailVo 结构体
type RoomDetailVo struct {
	// 价格分组
	PriceInfoGroups []RoomTypeBedInfoDto `json:"price_info_groups,omitempty" xml:"price_info_groups>room_type_bed_info_dto,omitempty"`
	// 图片
	Pics []string `json:"pics,omitempty" xml:"pics>string,omitempty"`
	// 房型设施分组
	FacilityGroupList []FacilityListVO `json:"facility_group_list,omitempty" xml:"facility_group_list>facility_list_vo,omitempty"`
	// 最低价含税  单位元
	LowestPriceToString string `json:"lowest_price_to_string,omitempty" xml:"lowest_price_to_string,omitempty"`
	// 货币类型
	PriceType string `json:"price_type,omitempty" xml:"price_type,omitempty"`
	// 房型名
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// 房型最低价含税 单位分
	LowestPrice int64 `json:"lowest_price,omitempty" xml:"lowest_price,omitempty"`
	// 是否加床
	AddBed int64 `json:"add_bed,omitempty" xml:"add_bed,omitempty"`
	// 房型id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// 床型信息
	RoomTypeBedInfo *RoomTypeBedInfoDto `json:"room_type_bed_info,omitempty" xml:"room_type_bed_info,omitempty"`
	// 房间属性
	RoomProperty *RoomProperty `json:"room_property,omitempty" xml:"room_property,omitempty"`
	// 满房
	Full bool `json:"full,omitempty" xml:"full,omitempty"`
}
