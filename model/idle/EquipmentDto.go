package idle

import (
	"sync"
)

// EquipmentDto 结构体
type EquipmentDto struct {
	// 标配名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 标配值，多为数量
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolEquipmentDto = sync.Pool{
	New: func() any {
		return new(EquipmentDto)
	},
}

// GetEquipmentDto() 从对象池中获取EquipmentDto
func GetEquipmentDto() *EquipmentDto {
	return poolEquipmentDto.Get().(*EquipmentDto)
}

// ReleaseEquipmentDto 释放EquipmentDto
func ReleaseEquipmentDto(v *EquipmentDto) {
	v.Name = ""
	v.Value = ""
	poolEquipmentDto.Put(v)
}
