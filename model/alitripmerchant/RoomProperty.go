package alitripmerchant

import (
	"sync"
)

// RoomProperty 结构体
type RoomProperty struct {
	// 面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 宽带
	NetworkService string `json:"network_service,omitempty" xml:"network_service,omitempty"`
	// 床型icon
	BedTypeIcon string `json:"bed_type_icon,omitempty" xml:"bed_type_icon,omitempty"`
	// 最大人数icon
	MaxOccupancyIcon string `json:"max_occupancy_icon,omitempty" xml:"max_occupancy_icon,omitempty"`
	// 楼层
	Floor string `json:"floor,omitempty" xml:"floor,omitempty"`
	// 楼层icon
	FloorIcon string `json:"floor_icon,omitempty" xml:"floor_icon,omitempty"`
	// 面积icon
	AreaIcon string `json:"area_icon,omitempty" xml:"area_icon,omitempty"`
	// 窗型icon
	WindowTypeIcon string `json:"window_type_icon,omitempty" xml:"window_type_icon,omitempty"`
	// 窗型
	WindowType string `json:"window_type,omitempty" xml:"window_type,omitempty"`
	// 最大人数
	MaxOccupancy int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
}

var poolRoomProperty = sync.Pool{
	New: func() any {
		return new(RoomProperty)
	},
}

// GetRoomProperty() 从对象池中获取RoomProperty
func GetRoomProperty() *RoomProperty {
	return poolRoomProperty.Get().(*RoomProperty)
}

// ReleaseRoomProperty 释放RoomProperty
func ReleaseRoomProperty(v *RoomProperty) {
	v.Area = ""
	v.NetworkService = ""
	v.BedTypeIcon = ""
	v.MaxOccupancyIcon = ""
	v.Floor = ""
	v.FloorIcon = ""
	v.AreaIcon = ""
	v.WindowTypeIcon = ""
	v.WindowType = ""
	v.MaxOccupancy = 0
	poolRoomProperty.Put(v)
}
