package alitripmerchant

import (
	"sync"
)

// DerbyVoucherIdentityQrcodeVo 结构体
type DerbyVoucherIdentityQrcodeVo struct {
	// base64
	Qrcode string `json:"qrcode,omitempty" xml:"qrcode,omitempty"`
}

var poolDerbyVoucherIdentityQrcodeVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherIdentityQrcodeVo)
	},
}

// GetDerbyVoucherIdentityQrcodeVo() 从对象池中获取DerbyVoucherIdentityQrcodeVo
func GetDerbyVoucherIdentityQrcodeVo() *DerbyVoucherIdentityQrcodeVo {
	return poolDerbyVoucherIdentityQrcodeVo.Get().(*DerbyVoucherIdentityQrcodeVo)
}

// ReleaseDerbyVoucherIdentityQrcodeVo 释放DerbyVoucherIdentityQrcodeVo
func ReleaseDerbyVoucherIdentityQrcodeVo(v *DerbyVoucherIdentityQrcodeVo) {
	v.Qrcode = ""
	poolDerbyVoucherIdentityQrcodeVo.Put(v)
}
