package alitripmerchant

import (
	"sync"
)

// RoomPropertiesDto 结构体
type RoomPropertiesDto struct {
	// 设施类型
	SubType string `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
	// 设施名称
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolRoomPropertiesDto = sync.Pool{
	New: func() any {
		return new(RoomPropertiesDto)
	},
}

// GetRoomPropertiesDto() 从对象池中获取RoomPropertiesDto
func GetRoomPropertiesDto() *RoomPropertiesDto {
	return poolRoomPropertiesDto.Get().(*RoomPropertiesDto)
}

// ReleaseRoomPropertiesDto 释放RoomPropertiesDto
func ReleaseRoomPropertiesDto(v *RoomPropertiesDto) {
	v.SubType = ""
	v.Value = ""
	poolRoomPropertiesDto.Put(v)
}
