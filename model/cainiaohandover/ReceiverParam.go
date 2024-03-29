package cainiaohandover

import (
	"sync"
)

// ReceiverParam 结构体
type ReceiverParam struct {
	// 收件人名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 收件人座机号
	Telephone string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 收件人手机号
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 用户昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 收件人地址信息
	AddressParam *OpenAddressParam `json:"address_param,omitempty" xml:"address_param,omitempty"`
}

var poolReceiverParam = sync.Pool{
	New: func() any {
		return new(ReceiverParam)
	},
}

// GetReceiverParam() 从对象池中获取ReceiverParam
func GetReceiverParam() *ReceiverParam {
	return poolReceiverParam.Get().(*ReceiverParam)
}

// ReleaseReceiverParam 释放ReceiverParam
func ReleaseReceiverParam(v *ReceiverParam) {
	v.Name = ""
	v.Telephone = ""
	v.MobilePhone = ""
	v.UserNick = ""
	v.Email = ""
	v.AddressParam = nil
	poolReceiverParam.Put(v)
}
