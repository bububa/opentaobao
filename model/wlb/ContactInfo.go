package wlb

import (
	"sync"
)

// ContactInfo 结构体
type ContactInfo struct {
	// 仓库联系人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 联系电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
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
	v.Name = ""
	v.Phone = ""
	poolContactInfo.Put(v)
}
