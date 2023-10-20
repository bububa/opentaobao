package xhotelonlineorder

import (
	"sync"
)

// XOrderGuest 结构体
type XOrderGuest struct {
	// 入住人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 房间序号
	RoomPos int64 `json:"room_pos,omitempty" xml:"room_pos,omitempty"`
	// 入住人序号
	PersonPos int64 `json:"person_pos,omitempty" xml:"person_pos,omitempty"`
}

var poolXOrderGuest = sync.Pool{
	New: func() any {
		return new(XOrderGuest)
	},
}

// GetXOrderGuest() 从对象池中获取XOrderGuest
func GetXOrderGuest() *XOrderGuest {
	return poolXOrderGuest.Get().(*XOrderGuest)
}

// ReleaseXOrderGuest 释放XOrderGuest
func ReleaseXOrderGuest(v *XOrderGuest) {
	v.Name = ""
	v.RoomPos = 0
	v.PersonPos = 0
	poolXOrderGuest.Put(v)
}
