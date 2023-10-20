package icbulogistics

import (
	"sync"
)

// ConsigneeAddress 结构体
type ConsigneeAddress struct {
	// 地址所有者邮箱(卖家维护收货地址, 值等于买家邮箱)
	AddressEmail string `json:"address_email,omitempty" xml:"address_email,omitempty"`
	// 公司英文名
	CompanyNameEn string `json:"company_name_en,omitempty" xml:"company_name_en,omitempty"`
	// 联系人姓名
	ContactPerson string `json:"contact_person,omitempty" xml:"contact_person,omitempty"`
	// 地址类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 公司中文名
	CompanyNameCn string `json:"company_name_cn,omitempty" xml:"company_name_cn,omitempty"`
	// 国家、省、市、详细地址信息
	Address *Address `json:"address,omitempty" xml:"address,omitempty"`
	// 联系方式(邮箱、电话号码、手机号码等)
	Contact *Contact `json:"contact,omitempty" xml:"contact,omitempty"`
}

var poolConsigneeAddress = sync.Pool{
	New: func() any {
		return new(ConsigneeAddress)
	},
}

// GetConsigneeAddress() 从对象池中获取ConsigneeAddress
func GetConsigneeAddress() *ConsigneeAddress {
	return poolConsigneeAddress.Get().(*ConsigneeAddress)
}

// ReleaseConsigneeAddress 释放ConsigneeAddress
func ReleaseConsigneeAddress(v *ConsigneeAddress) {
	v.AddressEmail = ""
	v.CompanyNameEn = ""
	v.ContactPerson = ""
	v.Type = ""
	v.CompanyNameCn = ""
	v.Address = nil
	v.Contact = nil
	poolConsigneeAddress.Put(v)
}
