package hotelhstdf

import (
	"sync"
)

// RoomTypePo 结构体
type RoomTypePo struct {
	// 外部系统id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 卖家未匹配房型的面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 卖家未匹配房型的床型
	BedType string `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	// 卖家未匹配房型英文名称
	NameE string `json:"name_e,omitempty" xml:"name_e,omitempty"`
	// 卖家未匹配房型名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 卖家未匹配rid
	Rid int64 `json:"rid,omitempty" xml:"rid,omitempty"`
	// 窗型
	WindowType int64 `json:"window_type,omitempty" xml:"window_type,omitempty"`
}

var poolRoomTypePo = sync.Pool{
	New: func() any {
		return new(RoomTypePo)
	},
}

// GetRoomTypePo() 从对象池中获取RoomTypePo
func GetRoomTypePo() *RoomTypePo {
	return poolRoomTypePo.Get().(*RoomTypePo)
}

// ReleaseRoomTypePo 释放RoomTypePo
func ReleaseRoomTypePo(v *RoomTypePo) {
	v.OuterId = ""
	v.Area = ""
	v.BedType = ""
	v.NameE = ""
	v.Name = ""
	v.Rid = 0
	v.WindowType = 0
	poolRoomTypePo.Put(v)
}
