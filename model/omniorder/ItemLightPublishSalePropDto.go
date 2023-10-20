package omniorder

import (
	"sync"
)

// ItemLightPublishSalePropDto 结构体
type ItemLightPublishSalePropDto struct {
	// value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// pid
	Pid int64 `json:"pid,omitempty" xml:"pid,omitempty"`
}

var poolItemLightPublishSalePropDto = sync.Pool{
	New: func() any {
		return new(ItemLightPublishSalePropDto)
	},
}

// GetItemLightPublishSalePropDto() 从对象池中获取ItemLightPublishSalePropDto
func GetItemLightPublishSalePropDto() *ItemLightPublishSalePropDto {
	return poolItemLightPublishSalePropDto.Get().(*ItemLightPublishSalePropDto)
}

// ReleaseItemLightPublishSalePropDto 释放ItemLightPublishSalePropDto
func ReleaseItemLightPublishSalePropDto(v *ItemLightPublishSalePropDto) {
	v.Value = ""
	v.Pid = 0
	poolItemLightPublishSalePropDto.Put(v)
}
