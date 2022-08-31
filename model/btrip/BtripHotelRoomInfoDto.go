package btrip

// BtripHotelRoomInfoDto 结构体
type BtripHotelRoomInfoDto struct {
	// 酒店房间设施列表
	RoomFacilityList []string `json:"room_facility_list,omitempty" xml:"room_facility_list>string,omitempty"`
	// 房型名称
	RoomTypeName string `json:"room_type_name,omitempty" xml:"room_type_name,omitempty"`
}
