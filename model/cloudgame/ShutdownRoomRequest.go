package cloudgame

import (
	"sync"
)

// ShutdownRoomRequest 结构体
type ShutdownRoomRequest struct {
	// 验签token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 房间ID
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
}

var poolShutdownRoomRequest = sync.Pool{
	New: func() any {
		return new(ShutdownRoomRequest)
	},
}

// GetShutdownRoomRequest() 从对象池中获取ShutdownRoomRequest
func GetShutdownRoomRequest() *ShutdownRoomRequest {
	return poolShutdownRoomRequest.Get().(*ShutdownRoomRequest)
}

// ReleaseShutdownRoomRequest 释放ShutdownRoomRequest
func ReleaseShutdownRoomRequest(v *ShutdownRoomRequest) {
	v.Token = ""
	v.RoomId = 0
	poolShutdownRoomRequest.Put(v)
}
