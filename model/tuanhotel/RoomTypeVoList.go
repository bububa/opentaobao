package tuanhotel

import (
	"sync"
)

// RoomTypeVoList 结构体
type RoomTypeVoList struct {
	// 是否无线上网
	NetworkService string `json:"network_service,omitempty" xml:"network_service,omitempty"`
	// 面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 床型
	Bed string `json:"bed,omitempty" xml:"bed,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 最多可入住人数
	MaxOccupancy string `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	// 楼层
	Floor string `json:"floor,omitempty" xml:"floor,omitempty"`
	// rId
	Rid string `json:"rid,omitempty" xml:"rid,omitempty"`
	// 设施
	Facility string `json:"facility,omitempty" xml:"facility,omitempty"`
	// 是否标准
	Standard bool `json:"standard,omitempty" xml:"standard,omitempty"`
}

var poolRoomTypeVoList = sync.Pool{
	New: func() any {
		return new(RoomTypeVoList)
	},
}

// GetRoomTypeVoList() 从对象池中获取RoomTypeVoList
func GetRoomTypeVoList() *RoomTypeVoList {
	return poolRoomTypeVoList.Get().(*RoomTypeVoList)
}

// ReleaseRoomTypeVoList 释放RoomTypeVoList
func ReleaseRoomTypeVoList(v *RoomTypeVoList) {
	v.NetworkService = ""
	v.Area = ""
	v.Bed = ""
	v.Name = ""
	v.MaxOccupancy = ""
	v.Floor = ""
	v.Rid = ""
	v.Facility = ""
	v.Standard = false
	poolRoomTypeVoList.Put(v)
}
