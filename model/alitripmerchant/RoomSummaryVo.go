package alitripmerchant

import (
	"sync"
)

// RoomSummaryVo 结构体
type RoomSummaryVo struct {
	// 房型名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 房型id
	SrId int64 `json:"sr_id,omitempty" xml:"sr_id,omitempty"`
}

var poolRoomSummaryVo = sync.Pool{
	New: func() any {
		return new(RoomSummaryVo)
	},
}

// GetRoomSummaryVo() 从对象池中获取RoomSummaryVo
func GetRoomSummaryVo() *RoomSummaryVo {
	return poolRoomSummaryVo.Get().(*RoomSummaryVo)
}

// ReleaseRoomSummaryVo 释放RoomSummaryVo
func ReleaseRoomSummaryVo(v *RoomSummaryVo) {
	v.Name = ""
	v.SrId = 0
	poolRoomSummaryVo.Put(v)
}
