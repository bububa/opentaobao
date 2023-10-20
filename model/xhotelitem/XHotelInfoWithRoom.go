package xhotelitem

import (
	"sync"
)

// XHotelInfoWithRoom 结构体
type XHotelInfoWithRoom struct {
	// 房型基础信息
	RoomTypeList []RoomType `json:"room_type_list,omitempty" xml:"room_type_list>room_type,omitempty"`
}

var poolXHotelInfoWithRoom = sync.Pool{
	New: func() any {
		return new(XHotelInfoWithRoom)
	},
}

// GetXHotelInfoWithRoom() 从对象池中获取XHotelInfoWithRoom
func GetXHotelInfoWithRoom() *XHotelInfoWithRoom {
	return poolXHotelInfoWithRoom.Get().(*XHotelInfoWithRoom)
}

// ReleaseXHotelInfoWithRoom 释放XHotelInfoWithRoom
func ReleaseXHotelInfoWithRoom(v *XHotelInfoWithRoom) {
	v.RoomTypeList = v.RoomTypeList[:0]
	poolXHotelInfoWithRoom.Put(v)
}
