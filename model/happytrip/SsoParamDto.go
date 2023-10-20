package happytrip

import (
	"sync"
)

// SsoParamDto 结构体
type SsoParamDto struct {
	// 免登验证ticket
	Ticket string `json:"ticket,omitempty" xml:"ticket,omitempty"`
}

var poolSsoParamDto = sync.Pool{
	New: func() any {
		return new(SsoParamDto)
	},
}

// GetSsoParamDto() 从对象池中获取SsoParamDto
func GetSsoParamDto() *SsoParamDto {
	return poolSsoParamDto.Get().(*SsoParamDto)
}

// ReleaseSsoParamDto 释放SsoParamDto
func ReleaseSsoParamDto(v *SsoParamDto) {
	v.Ticket = ""
	poolSsoParamDto.Put(v)
}
