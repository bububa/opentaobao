package alitripmerchant

import (
	"sync"
)

// DerbyVoucherVo 结构体
type DerbyVoucherVo struct {
	// 线下权益券二维码
	QRCodeImage string `json:"q_r_code_image,omitempty" xml:"q_r_code_image,omitempty"`
	// 线下核销是否成功
	OfflineRedemption bool `json:"offline_redemption,omitempty" xml:"offline_redemption,omitempty"`
}

var poolDerbyVoucherVo = sync.Pool{
	New: func() any {
		return new(DerbyVoucherVo)
	},
}

// GetDerbyVoucherVo() 从对象池中获取DerbyVoucherVo
func GetDerbyVoucherVo() *DerbyVoucherVo {
	return poolDerbyVoucherVo.Get().(*DerbyVoucherVo)
}

// ReleaseDerbyVoucherVo 释放DerbyVoucherVo
func ReleaseDerbyVoucherVo(v *DerbyVoucherVo) {
	v.QRCodeImage = ""
	v.OfflineRedemption = false
	poolDerbyVoucherVo.Put(v)
}
