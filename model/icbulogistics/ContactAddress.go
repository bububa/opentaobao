package icbulogistics

import (
	"sync"
)

// ContactAddress 结构体
type ContactAddress struct {
	// 联系人
	ContactPerson string `json:"contact_person,omitempty" xml:"contact_person,omitempty"`
	// 公司名称
	CompanyNameCn string `json:"company_name_cn,omitempty" xml:"company_name_cn,omitempty"`
	// 国家、省、市、详细地址信息
	Address *Address `json:"address,omitempty" xml:"address,omitempty"`
	// 联系方式(邮箱、电话号码、手机号码等)
	Contact *Contact `json:"contact,omitempty" xml:"contact,omitempty"`
}

var poolContactAddress = sync.Pool{
	New: func() any {
		return new(ContactAddress)
	},
}

// GetContactAddress() 从对象池中获取ContactAddress
func GetContactAddress() *ContactAddress {
	return poolContactAddress.Get().(*ContactAddress)
}

// ReleaseContactAddress 释放ContactAddress
func ReleaseContactAddress(v *ContactAddress) {
	v.ContactPerson = ""
	v.CompanyNameCn = ""
	v.Address = nil
	v.Contact = nil
	poolContactAddress.Put(v)
}
