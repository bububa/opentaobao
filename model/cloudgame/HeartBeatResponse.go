package cloudgame

import (
	"sync"
)

// HeartBeatResponse 结构体
type HeartBeatResponse struct {
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 扩展字段
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 房间id
	RoomId int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// 上报时间戳
	ReportTime int64 `json:"report_time,omitempty" xml:"report_time,omitempty"`
	// 上报序列
	Seq int64 `json:"seq,omitempty" xml:"seq,omitempty"`
}

var poolHeartBeatResponse = sync.Pool{
	New: func() any {
		return new(HeartBeatResponse)
	},
}

// GetHeartBeatResponse() 从对象池中获取HeartBeatResponse
func GetHeartBeatResponse() *HeartBeatResponse {
	return poolHeartBeatResponse.Get().(*HeartBeatResponse)
}

// ReleaseHeartBeatResponse 释放HeartBeatResponse
func ReleaseHeartBeatResponse(v *HeartBeatResponse) {
	v.UserId = ""
	v.ExtInfo = ""
	v.RoomId = 0
	v.ReportTime = 0
	v.Seq = 0
	poolHeartBeatResponse.Put(v)
}
