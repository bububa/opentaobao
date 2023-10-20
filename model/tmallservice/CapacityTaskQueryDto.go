package tmallservice

import (
	"sync"
)

// CapacityTaskQueryDto 结构体
type CapacityTaskQueryDto struct {
	// 省份
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
}

var poolCapacityTaskQueryDto = sync.Pool{
	New: func() any {
		return new(CapacityTaskQueryDto)
	},
}

// GetCapacityTaskQueryDto() 从对象池中获取CapacityTaskQueryDto
func GetCapacityTaskQueryDto() *CapacityTaskQueryDto {
	return poolCapacityTaskQueryDto.Get().(*CapacityTaskQueryDto)
}

// ReleaseCapacityTaskQueryDto 释放CapacityTaskQueryDto
func ReleaseCapacityTaskQueryDto(v *CapacityTaskQueryDto) {
	v.ProvinceName = ""
	poolCapacityTaskQueryDto.Put(v)
}
