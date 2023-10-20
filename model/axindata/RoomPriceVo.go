package axindata

import (
	"sync"
)

// RoomPriceVo 结构体
type RoomPriceVo struct {
	// 价库信息
	RateList []RateVo `json:"rate_list,omitempty" xml:"rate_list>rate_vo,omitempty"`
	// 房型信息
	StdRoomInfo *StdRoomType `json:"std_room_info,omitempty" xml:"std_room_info,omitempty"`
}

var poolRoomPriceVo = sync.Pool{
	New: func() any {
		return new(RoomPriceVo)
	},
}

// GetRoomPriceVo() 从对象池中获取RoomPriceVo
func GetRoomPriceVo() *RoomPriceVo {
	return poolRoomPriceVo.Get().(*RoomPriceVo)
}

// ReleaseRoomPriceVo 释放RoomPriceVo
func ReleaseRoomPriceVo(v *RoomPriceVo) {
	v.RateList = v.RateList[:0]
	v.StdRoomInfo = nil
	poolRoomPriceVo.Put(v)
}
