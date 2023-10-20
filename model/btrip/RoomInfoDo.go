package btrip

import (
	"sync"
)

// RoomInfoDo 结构体
type RoomInfoDo struct {
	// 房间名称
	RoomName string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// 货币种类
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 房间价格（分）
	RoomPrice int64 `json:"room_price,omitempty" xml:"room_price,omitempty"`
	// 房间数量
	RoomNum int64 `json:"room_num,omitempty" xml:"room_num,omitempty"`
	// 总间夜
	NightNumber int64 `json:"night_number,omitempty" xml:"night_number,omitempty"`
}

var poolRoomInfoDo = sync.Pool{
	New: func() any {
		return new(RoomInfoDo)
	},
}

// GetRoomInfoDo() 从对象池中获取RoomInfoDo
func GetRoomInfoDo() *RoomInfoDo {
	return poolRoomInfoDo.Get().(*RoomInfoDo)
}

// ReleaseRoomInfoDo 释放RoomInfoDo
func ReleaseRoomInfoDo(v *RoomInfoDo) {
	v.RoomName = ""
	v.Currency = ""
	v.RoomPrice = 0
	v.RoomNum = 0
	v.NightNumber = 0
	poolRoomInfoDo.Put(v)
}
