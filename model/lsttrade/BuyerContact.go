package lsttrade

import (
	"sync"
)

// BuyerContact 结构体
type BuyerContact struct {
	// 买家电话
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

var poolBuyerContact = sync.Pool{
	New: func() any {
		return new(BuyerContact)
	},
}

// GetBuyerContact() 从对象池中获取BuyerContact
func GetBuyerContact() *BuyerContact {
	return poolBuyerContact.Get().(*BuyerContact)
}

// ReleaseBuyerContact 释放BuyerContact
func ReleaseBuyerContact(v *BuyerContact) {
	v.Phone = ""
	v.CompanyName = ""
	v.Mobile = ""
	v.Name = ""
	v.Email = ""
	poolBuyerContact.Put(v)
}
