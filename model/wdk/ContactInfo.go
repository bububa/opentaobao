package wdk

import (
	"sync"
)

// ContactInfo 结构体
type ContactInfo struct {
	// 联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人类型
	ContactType string `json:"contact_type,omitempty" xml:"contact_type,omitempty"`
	// 联系人手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

var poolContactInfo = sync.Pool{
	New: func() any {
		return new(ContactInfo)
	},
}

// GetContactInfo() 从对象池中获取ContactInfo
func GetContactInfo() *ContactInfo {
	return poolContactInfo.Get().(*ContactInfo)
}

// ReleaseContactInfo 释放ContactInfo
func ReleaseContactInfo(v *ContactInfo) {
	v.ContactName = ""
	v.ContactType = ""
	v.Mobile = ""
	poolContactInfo.Put(v)
}
