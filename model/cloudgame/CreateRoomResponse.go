package cloudgame

import (
	"sync"
)

// CreateRoomResponse 结构体
type CreateRoomResponse struct {
	// 扩展数据
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}

var poolCreateRoomResponse = sync.Pool{
	New: func() any {
		return new(CreateRoomResponse)
	},
}

// GetCreateRoomResponse() 从对象池中获取CreateRoomResponse
func GetCreateRoomResponse() *CreateRoomResponse {
	return poolCreateRoomResponse.Get().(*CreateRoomResponse)
}

// ReleaseCreateRoomResponse 释放CreateRoomResponse
func ReleaseCreateRoomResponse(v *CreateRoomResponse) {
	v.ExtInfo = ""
	v.RoomId = 0
	poolCreateRoomResponse.Put(v)
}
