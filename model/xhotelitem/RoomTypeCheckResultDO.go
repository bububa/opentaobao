package xhotelitem

// RoomTypeCheckResultDo 结构体
type RoomTypeCheckResultDo struct {
	// 床型数据检查结果，false代表有冲突
	Bedtypecheckresult string `json:"bedtypecheckresult,omitempty" xml:"bedtypecheckresult,omitempty"`
	// 结构化床型
	BedType string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	// 原始床型
	OriginBedType string `json:"origin_bed_type,omitempty" xml:"origin_bed_type,omitempty"`
	// 房型英文名称
	Roomtypeenname string `json:"roomtypeenname,omitempty" xml:"roomtypeenname,omitempty"`
	// 房型名称
	Roomtypename string `json:"roomtypename,omitempty" xml:"roomtypename,omitempty"`
	// 房型编码
	OuterRoomTypeId string `json:"outer_room_type_id,omitempty" xml:"outer_room_type_id,omitempty"`
	// 酒店名称
	Hotelname string `json:"hotelname,omitempty" xml:"hotelname,omitempty"`
	// 酒店编码
	OuterHotelId string `json:"outer_hotel_id,omitempty" xml:"outer_hotel_id,omitempty"`
	// 冲突数据创建时间（供商家参考，如果时间是过去三天以外的，那么可以不用关注）
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 床型英文描叙
	Ennamebedtypedesc string `json:"ennamebedtypedesc,omitempty" xml:"ennamebedtypedesc,omitempty"`
	// 床型描叙
	Namebedtypedesc string `json:"namebedtypedesc,omitempty" xml:"namebedtypedesc,omitempty"`
	// 床型描叙
	Bedtypedesc string `json:"bedtypedesc,omitempty" xml:"bedtypedesc,omitempty"`
	// 床型检查信息
	Bedtypecheckmsg string `json:"bedtypecheckmsg,omitempty" xml:"bedtypecheckmsg,omitempty"`
}
