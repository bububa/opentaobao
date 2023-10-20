package hotel

import (
	"sync"
)

// SRoomTypePropertiesSetVo 结构体
type SRoomTypePropertiesSetVo struct {
	// 标准房型普通图集合
	SroomTypeNomalPictures []ShotelPropertiesVo `json:"sroom_type_nomal_pictures,omitempty" xml:"sroom_type_nomal_pictures>shotel_properties_vo,omitempty"`
	// 房间设施集合
	SroomTypeRoomFacilities []ShotelPropertiesVo `json:"sroom_type_room_facilities,omitempty" xml:"sroom_type_room_facilities>shotel_properties_vo,omitempty"`
}

var poolSRoomTypePropertiesSetVo = sync.Pool{
	New: func() any {
		return new(SRoomTypePropertiesSetVo)
	},
}

// GetSRoomTypePropertiesSetVo() 从对象池中获取SRoomTypePropertiesSetVo
func GetSRoomTypePropertiesSetVo() *SRoomTypePropertiesSetVo {
	return poolSRoomTypePropertiesSetVo.Get().(*SRoomTypePropertiesSetVo)
}

// ReleaseSRoomTypePropertiesSetVo 释放SRoomTypePropertiesSetVo
func ReleaseSRoomTypePropertiesSetVo(v *SRoomTypePropertiesSetVo) {
	v.SroomTypeNomalPictures = v.SroomTypeNomalPictures[:0]
	v.SroomTypeRoomFacilities = v.SroomTypeRoomFacilities[:0]
	poolSRoomTypePropertiesSetVo.Put(v)
}
