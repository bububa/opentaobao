package lstlogistics2

import (
	"sync"
)

// ContactParam 结构体
type ContactParam struct {
	// 联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 手机号为11位
	ContactMobile string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
	// 电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
}

var poolContactParam = sync.Pool{
	New: func() any {
		return new(ContactParam)
	},
}

// GetContactParam() 从对象池中获取ContactParam
func GetContactParam() *ContactParam {
	return poolContactParam.Get().(*ContactParam)
}

// ReleaseContactParam 释放ContactParam
func ReleaseContactParam(v *ContactParam) {
	v.ContactName = ""
	v.ContactMobile = ""
	v.ContactPhone = ""
	poolContactParam.Put(v)
}
