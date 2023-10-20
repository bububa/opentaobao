package ieagency

import (
	"sync"
)

// ContactsParam 结构体
type ContactsParam struct {
	// 联系人邮箱地址
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 联系人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 联系人手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 联系人手机号国家编码
	PhoneCountryCode string `json:"phone_country_code,omitempty" xml:"phone_country_code,omitempty"`
}

var poolContactsParam = sync.Pool{
	New: func() any {
		return new(ContactsParam)
	},
}

// GetContactsParam() 从对象池中获取ContactsParam
func GetContactsParam() *ContactsParam {
	return poolContactsParam.Get().(*ContactsParam)
}

// ReleaseContactsParam 释放ContactsParam
func ReleaseContactsParam(v *ContactsParam) {
	v.Email = ""
	v.Name = ""
	v.Phone = ""
	v.PhoneCountryCode = ""
	poolContactsParam.Put(v)
}
