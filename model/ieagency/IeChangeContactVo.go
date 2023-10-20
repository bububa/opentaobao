package ieagency

import (
	"sync"
)

// IeChangeContactVo 结构体
type IeChangeContactVo struct {
	// 电话国际区号
	MobileCountryCode string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// 电话号
	MobilePhoneNumber string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolIeChangeContactVo = sync.Pool{
	New: func() any {
		return new(IeChangeContactVo)
	},
}

// GetIeChangeContactVo() 从对象池中获取IeChangeContactVo
func GetIeChangeContactVo() *IeChangeContactVo {
	return poolIeChangeContactVo.Get().(*IeChangeContactVo)
}

// ReleaseIeChangeContactVo 释放IeChangeContactVo
func ReleaseIeChangeContactVo(v *IeChangeContactVo) {
	v.MobileCountryCode = ""
	v.MobilePhoneNumber = ""
	v.Email = ""
	v.Name = ""
	poolIeChangeContactVo.Put(v)
}
