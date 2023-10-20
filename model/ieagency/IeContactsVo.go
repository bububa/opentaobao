package ieagency

import (
	"sync"
)

// IeContactsVo 结构体
type IeContactsVo struct {
	// 联系人邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 联系人姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 联系人电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}

var poolIeContactsVo = sync.Pool{
	New: func() any {
		return new(IeContactsVo)
	},
}

// GetIeContactsVo() 从对象池中获取IeContactsVo
func GetIeContactsVo() *IeContactsVo {
	return poolIeContactsVo.Get().(*IeContactsVo)
}

// ReleaseIeContactsVo 释放IeContactsVo
func ReleaseIeContactsVo(v *IeContactsVo) {
	v.Email = ""
	v.Name = ""
	v.Phone = ""
	poolIeContactsVo.Put(v)
}
