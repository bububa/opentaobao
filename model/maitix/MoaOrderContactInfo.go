package maitix

import (
	"sync"
)

// MoaOrderContactInfo 结构体
type MoaOrderContactInfo struct {
	// 联系人姓名-必填
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人国家码-必填
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 联系人email
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 联系人手机号-必填
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}

var poolMoaOrderContactInfo = sync.Pool{
	New: func() any {
		return new(MoaOrderContactInfo)
	},
}

// GetMoaOrderContactInfo() 从对象池中获取MoaOrderContactInfo
func GetMoaOrderContactInfo() *MoaOrderContactInfo {
	return poolMoaOrderContactInfo.Get().(*MoaOrderContactInfo)
}

// ReleaseMoaOrderContactInfo 释放MoaOrderContactInfo
func ReleaseMoaOrderContactInfo(v *MoaOrderContactInfo) {
	v.ContactName = ""
	v.CountryCode = ""
	v.Email = ""
	v.Phone = ""
	poolMoaOrderContactInfo.Put(v)
}
