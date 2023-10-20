package nazca

import (
	"sync"
)

// ChangeAuthCallBackDo 结构体
type ChangeAuthCallBackDo struct {
	// 身份证号
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 支付宝账号
	AlipayAccount string `json:"alipay_account,omitempty" xml:"alipay_account,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 姓名
	PersonName string `json:"person_name,omitempty" xml:"person_name,omitempty"`
	// 客户在1688的唯一标识
	PlatformUserId string `json:"platform_user_id,omitempty" xml:"platform_user_id,omitempty"`
	// 认证状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 联系手机
	Telphone string `json:"telphone,omitempty" xml:"telphone,omitempty"`
}

var poolChangeAuthCallBackDo = sync.Pool{
	New: func() any {
		return new(ChangeAuthCallBackDo)
	},
}

// GetChangeAuthCallBackDo() 从对象池中获取ChangeAuthCallBackDo
func GetChangeAuthCallBackDo() *ChangeAuthCallBackDo {
	return poolChangeAuthCallBackDo.Get().(*ChangeAuthCallBackDo)
}

// ReleaseChangeAuthCallBackDo 释放ChangeAuthCallBackDo
func ReleaseChangeAuthCallBackDo(v *ChangeAuthCallBackDo) {
	v.Id = ""
	v.AlipayAccount = ""
	v.Email = ""
	v.PersonName = ""
	v.PlatformUserId = ""
	v.Status = ""
	v.Telphone = ""
	poolChangeAuthCallBackDo.Put(v)
}
