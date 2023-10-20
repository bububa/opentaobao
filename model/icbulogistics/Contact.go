package icbulogistics

import (
	"sync"
)

// Contact 结构体
type Contact struct {
	// 电话区号
	PhoneCode string `json:"phone_code,omitempty" xml:"phone_code,omitempty"`
	// 手机号码
	MobileNo string `json:"mobile_no,omitempty" xml:"mobile_no,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 电话区码
	PhoneArea string `json:"phone_area,omitempty" xml:"phone_area,omitempty"`
}

var poolContact = sync.Pool{
	New: func() any {
		return new(Contact)
	},
}

// GetContact() 从对象池中获取Contact
func GetContact() *Contact {
	return poolContact.Get().(*Contact)
}

// ReleaseContact 释放Contact
func ReleaseContact(v *Contact) {
	v.PhoneCode = ""
	v.MobileNo = ""
	v.Email = ""
	v.PhoneArea = ""
	poolContact.Put(v)
}
