package flightuppc

import (
	"sync"
)

// AlipayDecryptReq 结构体
type AlipayDecryptReq struct {
	// 待解密密文
	EncryptContent string `json:"encrypt_content,omitempty" xml:"encrypt_content,omitempty"`
	// 业务代码
	ExternalAppletBizCode string `json:"external_applet_biz_code,omitempty" xml:"external_applet_biz_code,omitempty"`
}

var poolAlipayDecryptReq = sync.Pool{
	New: func() any {
		return new(AlipayDecryptReq)
	},
}

// GetAlipayDecryptReq() 从对象池中获取AlipayDecryptReq
func GetAlipayDecryptReq() *AlipayDecryptReq {
	return poolAlipayDecryptReq.Get().(*AlipayDecryptReq)
}

// ReleaseAlipayDecryptReq 释放AlipayDecryptReq
func ReleaseAlipayDecryptReq(v *AlipayDecryptReq) {
	v.EncryptContent = ""
	v.ExternalAppletBizCode = ""
	poolAlipayDecryptReq.Put(v)
}
