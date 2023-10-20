package trade

import (
	"sync"
)

// OfnSelfRecycleAuthDto 结构体
type OfnSelfRecycleAuthDto struct {
	// 是否鉴权通过
	AuthPass bool `json:"auth_pass,omitempty" xml:"auth_pass,omitempty"`
}

var poolOfnSelfRecycleAuthDto = sync.Pool{
	New: func() any {
		return new(OfnSelfRecycleAuthDto)
	},
}

// GetOfnSelfRecycleAuthDto() 从对象池中获取OfnSelfRecycleAuthDto
func GetOfnSelfRecycleAuthDto() *OfnSelfRecycleAuthDto {
	return poolOfnSelfRecycleAuthDto.Get().(*OfnSelfRecycleAuthDto)
}

// ReleaseOfnSelfRecycleAuthDto 释放OfnSelfRecycleAuthDto
func ReleaseOfnSelfRecycleAuthDto(v *OfnSelfRecycleAuthDto) {
	v.AuthPass = false
	poolOfnSelfRecycleAuthDto.Put(v)
}
