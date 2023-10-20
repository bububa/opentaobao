package btrip

import (
	"sync"
)

// ContactInfoDto 结构体
type ContactInfoDto struct {
	// 联系优先
	ContactEmail string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	// 联系人名字
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系手机号
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
}

var poolContactInfoDto = sync.Pool{
	New: func() any {
		return new(ContactInfoDto)
	},
}

// GetContactInfoDto() 从对象池中获取ContactInfoDto
func GetContactInfoDto() *ContactInfoDto {
	return poolContactInfoDto.Get().(*ContactInfoDto)
}

// ReleaseContactInfoDto 释放ContactInfoDto
func ReleaseContactInfoDto(v *ContactInfoDto) {
	v.ContactEmail = ""
	v.ContactName = ""
	v.ContactPhone = ""
	poolContactInfoDto.Put(v)
}
