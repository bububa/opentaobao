package feedflow

import (
	"sync"
)

// AccessAllowedDto 结构体
type AccessAllowedDto struct {
	// 不可以使用的原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 是否可以使用，false不可以进行广告投放
	IsAccessAllowed bool `json:"is_access_allowed,omitempty" xml:"is_access_allowed,omitempty"`
}

var poolAccessAllowedDto = sync.Pool{
	New: func() any {
		return new(AccessAllowedDto)
	},
}

// GetAccessAllowedDto() 从对象池中获取AccessAllowedDto
func GetAccessAllowedDto() *AccessAllowedDto {
	return poolAccessAllowedDto.Get().(*AccessAllowedDto)
}

// ReleaseAccessAllowedDto 释放AccessAllowedDto
func ReleaseAccessAllowedDto(v *AccessAllowedDto) {
	v.Reason = ""
	v.IsAccessAllowed = false
	poolAccessAllowedDto.Put(v)
}
