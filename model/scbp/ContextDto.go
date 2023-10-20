package scbp

import (
	"sync"
)

// ContextDto 结构体
type ContextDto struct {
	// login_id
	LoginId string `json:"login_id,omitempty" xml:"login_id,omitempty"`
	// from
	From string `json:"from,omitempty" xml:"from,omitempty"`
	// service_type
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// ip
	Ip string `json:"ip,omitempty" xml:"ip,omitempty"`
	// is_admin
	IsAdmin bool `json:"is_admin,omitempty" xml:"is_admin,omitempty"`
	// is_top
	IsTop bool `json:"is_top,omitempty" xml:"is_top,omitempty"`
}

var poolContextDto = sync.Pool{
	New: func() any {
		return new(ContextDto)
	},
}

// GetContextDto() 从对象池中获取ContextDto
func GetContextDto() *ContextDto {
	return poolContextDto.Get().(*ContextDto)
}

// ReleaseContextDto 释放ContextDto
func ReleaseContextDto(v *ContextDto) {
	v.LoginId = ""
	v.From = ""
	v.ServiceType = ""
	v.Ip = ""
	v.IsAdmin = false
	v.IsTop = false
	poolContextDto.Put(v)
}
