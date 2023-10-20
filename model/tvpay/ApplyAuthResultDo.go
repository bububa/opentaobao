package tvpay

import (
	"sync"
)

// ApplyAuthResultDo 结构体
type ApplyAuthResultDo struct {
	// 授权方式
	AuthMode string `json:"auth_mode,omitempty" xml:"auth_mode,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// qrcode
	QrCode string `json:"qr_code,omitempty" xml:"qr_code,omitempty"`
	// 二维码地址
	QrCodeUrl string `json:"qr_code_url,omitempty" xml:"qr_code_url,omitempty"`
}

var poolApplyAuthResultDo = sync.Pool{
	New: func() any {
		return new(ApplyAuthResultDo)
	},
}

// GetApplyAuthResultDo() 从对象池中获取ApplyAuthResultDo
func GetApplyAuthResultDo() *ApplyAuthResultDo {
	return poolApplyAuthResultDo.Get().(*ApplyAuthResultDo)
}

// ReleaseApplyAuthResultDo 释放ApplyAuthResultDo
func ReleaseApplyAuthResultDo(v *ApplyAuthResultDo) {
	v.AuthMode = ""
	v.Mobile = ""
	v.QrCode = ""
	v.QrCodeUrl = ""
	poolApplyAuthResultDo.Put(v)
}
