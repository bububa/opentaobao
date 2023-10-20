package axintrade

import (
	"sync"
)

// ContactInfoApiDto 结构体
type ContactInfoApiDto struct {
	// 联系人邮箱
	ContactEmail string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	// 联系人手机号
	ContactMobile string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
	// 联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人固定电话
	ContactTel string `json:"contact_tel,omitempty" xml:"contact_tel,omitempty"`
}

var poolContactInfoApiDto = sync.Pool{
	New: func() any {
		return new(ContactInfoApiDto)
	},
}

// GetContactInfoApiDto() 从对象池中获取ContactInfoApiDto
func GetContactInfoApiDto() *ContactInfoApiDto {
	return poolContactInfoApiDto.Get().(*ContactInfoApiDto)
}

// ReleaseContactInfoApiDto 释放ContactInfoApiDto
func ReleaseContactInfoApiDto(v *ContactInfoApiDto) {
	v.ContactEmail = ""
	v.ContactMobile = ""
	v.ContactName = ""
	v.ContactTel = ""
	poolContactInfoApiDto.Put(v)
}
