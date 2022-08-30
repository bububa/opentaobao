package axintrade

// RoomInfo 结构体
type RoomInfo struct {
	// 房型设施
	RoomFacilityList []string `json:"room_facility_list,omitempty" xml:"room_facility_list>string,omitempty"`
	// 房型名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
}
