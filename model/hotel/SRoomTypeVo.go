package hotel

// SRoomTypeVo 结构体
type SRoomTypeVo struct {
	// 是否可加床，0--不可，1--可以
	AddBed int64 `json:"add_bed,omitempty" xml:"add_bed,omitempty"`
	// 面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 设施文本
	Facility string `json:"facility,omitempty" xml:"facility,omitempty"`
	// 所在楼层
	Floor string `json:"floor,omitempty" xml:"floor,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 最后修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 最大入住人数
	MaxOccupancy int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	// 房型中文名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 房型英文名
	NameE string `json:"name_e,omitempty" xml:"name_e,omitempty"`
	// 床型结构化信息
	RoomTypeBedInfo *RoomTypeBedInfoVo `json:"room_type_bed_info,omitempty" xml:"room_type_bed_info,omitempty"`
	// 所属标准酒店id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 标准房型id
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
	// 上下架状态，0--下架，其他状态-下架
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 窗型，0--无窗,1--有窗,2--部分有窗,3--暗窗,4--部分暗窗
	WindowType string `json:"window_type,omitempty" xml:"window_type,omitempty"`
	// 房型图片，多值使用英文逗号拼接
	PicUrls string `json:"pic_urls,omitempty" xml:"pic_urls,omitempty"`
	// 房型附加属性集合
	SroomTypeProperties *SRoomTypePropertiesSetVo `json:"sroom_type_properties,omitempty" xml:"sroom_type_properties,omitempty"`
}
