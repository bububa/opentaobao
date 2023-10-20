package lsttrade

import (
	"sync"
)

// SellerContact 结构体
type SellerContact struct {
	// 电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 移动电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 邮件地址
	Email string `json:"email,omitempty" xml:"email,omitempty"`
}

var poolSellerContact = sync.Pool{
	New: func() any {
		return new(SellerContact)
	},
}

// GetSellerContact() 从对象池中获取SellerContact
func GetSellerContact() *SellerContact {
	return poolSellerContact.Get().(*SellerContact)
}

// ReleaseSellerContact 释放SellerContact
func ReleaseSellerContact(v *SellerContact) {
	v.Phone = ""
	v.CompanyName = ""
	v.Mobile = ""
	v.Name = ""
	v.Email = ""
	poolSellerContact.Put(v)
}
