package tblogistics

import (
	"sync"
)

// GenPickupCodeConfigTopDto 结构体
type GenPickupCodeConfigTopDto struct {
	// 取件码分组，分组下唯一
	Group string `json:"group,omitempty" xml:"group,omitempty"`
}

var poolGenPickupCodeConfigTopDto = sync.Pool{
	New: func() any {
		return new(GenPickupCodeConfigTopDto)
	},
}

// GetGenPickupCodeConfigTopDto() 从对象池中获取GenPickupCodeConfigTopDto
func GetGenPickupCodeConfigTopDto() *GenPickupCodeConfigTopDto {
	return poolGenPickupCodeConfigTopDto.Get().(*GenPickupCodeConfigTopDto)
}

// ReleaseGenPickupCodeConfigTopDto 释放GenPickupCodeConfigTopDto
func ReleaseGenPickupCodeConfigTopDto(v *GenPickupCodeConfigTopDto) {
	v.Group = ""
	poolGenPickupCodeConfigTopDto.Put(v)
}
