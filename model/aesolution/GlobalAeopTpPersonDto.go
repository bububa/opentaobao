package aesolution

import (
	"sync"
)

// GlobalAeopTpPersonDto 结构体
type GlobalAeopTpPersonDto struct {
	// login ID
	LoginId string `json:"login_id,omitempty" xml:"login_id,omitempty"`
	// last name
	LastName string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// first name
	FirstName string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// country/region
	Country string `json:"country,omitempty" xml:"country,omitempty"`
}

var poolGlobalAeopTpPersonDto = sync.Pool{
	New: func() any {
		return new(GlobalAeopTpPersonDto)
	},
}

// GetGlobalAeopTpPersonDto() 从对象池中获取GlobalAeopTpPersonDto
func GetGlobalAeopTpPersonDto() *GlobalAeopTpPersonDto {
	return poolGlobalAeopTpPersonDto.Get().(*GlobalAeopTpPersonDto)
}

// ReleaseGlobalAeopTpPersonDto 释放GlobalAeopTpPersonDto
func ReleaseGlobalAeopTpPersonDto(v *GlobalAeopTpPersonDto) {
	v.LoginId = ""
	v.LastName = ""
	v.FirstName = ""
	v.Country = ""
	poolGlobalAeopTpPersonDto.Put(v)
}
