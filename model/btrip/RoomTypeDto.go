package btrip

import (
	"sync"
)

// RoomTypeDto 结构体
type RoomTypeDto struct {
	// 面积
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 床型
	Bed string `json:"bed,omitempty" xml:"bed,omitempty"`
	// 设施
	Facility string `json:"facility,omitempty" xml:"facility,omitempty"`
	// 楼层
	Floor string `json:"floor,omitempty" xml:"floor,omitempty"`
	// 宽带类型
	Internet string `json:"internet,omitempty" xml:"internet,omitempty"`
	// 图片地址
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 窗型
	WindowType string `json:"window_type,omitempty" xml:"window_type,omitempty"`
	// 房型名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 最大入住人数
	MaxOccupancy int64 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	// 酒店Id
	Shid int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// 房型Id
	Srid int64 `json:"srid,omitempty" xml:"srid,omitempty"`
}

var poolRoomTypeDto = sync.Pool{
	New: func() any {
		return new(RoomTypeDto)
	},
}

// GetRoomTypeDto() 从对象池中获取RoomTypeDto
func GetRoomTypeDto() *RoomTypeDto {
	return poolRoomTypeDto.Get().(*RoomTypeDto)
}

// ReleaseRoomTypeDto 释放RoomTypeDto
func ReleaseRoomTypeDto(v *RoomTypeDto) {
	v.Area = ""
	v.Bed = ""
	v.Facility = ""
	v.Floor = ""
	v.Internet = ""
	v.PicUrl = ""
	v.WindowType = ""
	v.Name = ""
	v.MaxOccupancy = 0
	v.Shid = 0
	v.Srid = 0
	poolRoomTypeDto.Put(v)
}
