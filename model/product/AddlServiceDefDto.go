package product

import (
	"sync"
)

// AddlServiceDefDto 结构体
type AddlServiceDefDto struct {
	// 是否支持议价
	Bargain bool `json:"bargain,omitempty" xml:"bargain,omitempty"`
}

var poolAddlServiceDefDto = sync.Pool{
	New: func() any {
		return new(AddlServiceDefDto)
	},
}

// GetAddlServiceDefDto() 从对象池中获取AddlServiceDefDto
func GetAddlServiceDefDto() *AddlServiceDefDto {
	return poolAddlServiceDefDto.Get().(*AddlServiceDefDto)
}

// ReleaseAddlServiceDefDto 释放AddlServiceDefDto
func ReleaseAddlServiceDefDto(v *AddlServiceDefDto) {
	v.Bargain = false
	poolAddlServiceDefDto.Put(v)
}
