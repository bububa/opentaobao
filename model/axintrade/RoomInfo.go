package axintrade

import (
	"sync"
)

// RoomInfo 结构体
type RoomInfo struct {
	// 房型设施
	RoomFacilityList []string `json:"room_facility_list,omitempty" xml:"room_facility_list>string,omitempty"`
	// 房型名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
}

var poolRoomInfo = sync.Pool{
	New: func() any {
		return new(RoomInfo)
	},
}

// GetRoomInfo() 从对象池中获取RoomInfo
func GetRoomInfo() *RoomInfo {
	return poolRoomInfo.Get().(*RoomInfo)
}

// ReleaseRoomInfo 释放RoomInfo
func ReleaseRoomInfo(v *RoomInfo) {
	v.RoomFacilityList = v.RoomFacilityList[:0]
	v.RoomName = ""
	poolRoomInfo.Put(v)
}
