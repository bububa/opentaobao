package alitripmerchant

import (
	"sync"
)

// RegisterProgressVo 结构体
type RegisterProgressVo struct {
	// 在雅高官网注册时的手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 在雅高官网注册时的邮箱
	Email string `json:"email,omitempty" xml:"email,omitempty"`
	// 是否需要弹验证码弹框
	PopUpVerification bool `json:"pop_up_verification,omitempty" xml:"pop_up_verification,omitempty"`
	// 注册状态
	RegisterStatus bool `json:"register_status,omitempty" xml:"register_status,omitempty"`
}

var poolRegisterProgressVo = sync.Pool{
	New: func() any {
		return new(RegisterProgressVo)
	},
}

// GetRegisterProgressVo() 从对象池中获取RegisterProgressVo
func GetRegisterProgressVo() *RegisterProgressVo {
	return poolRegisterProgressVo.Get().(*RegisterProgressVo)
}

// ReleaseRegisterProgressVo 释放RegisterProgressVo
func ReleaseRegisterProgressVo(v *RegisterProgressVo) {
	v.Phone = ""
	v.Email = ""
	v.PopUpVerification = false
	v.RegisterStatus = false
	poolRegisterProgressVo.Put(v)
}
