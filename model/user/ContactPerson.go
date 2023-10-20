package user

import (
	"sync"
)

// ContactPerson 结构体
type ContactPerson struct {
	// 联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系电话
	ContactNumber string `json:"contact_number,omitempty" xml:"contact_number,omitempty"`
	// 联系人岗位
	ContactPosition string `json:"contact_position,omitempty" xml:"contact_position,omitempty"`
}

var poolContactPerson = sync.Pool{
	New: func() any {
		return new(ContactPerson)
	},
}

// GetContactPerson() 从对象池中获取ContactPerson
func GetContactPerson() *ContactPerson {
	return poolContactPerson.Get().(*ContactPerson)
}

// ReleaseContactPerson 释放ContactPerson
func ReleaseContactPerson(v *ContactPerson) {
	v.ContactName = ""
	v.ContactNumber = ""
	v.ContactPosition = ""
	poolContactPerson.Put(v)
}
