package flightuppc

import (
	"sync"
)

// AlipayEncryptReq 结构体
type AlipayEncryptReq struct {
	// 待加密明文
	SourceContent string `json:"source_content,omitempty" xml:"source_content,omitempty"`
	// 业务代码
	ExternalAppletBizCode string `json:"external_applet_biz_code,omitempty" xml:"external_applet_biz_code,omitempty"`
}

var poolAlipayEncryptReq = sync.Pool{
	New: func() any {
		return new(AlipayEncryptReq)
	},
}

// GetAlipayEncryptReq() 从对象池中获取AlipayEncryptReq
func GetAlipayEncryptReq() *AlipayEncryptReq {
	return poolAlipayEncryptReq.Get().(*AlipayEncryptReq)
}

// ReleaseAlipayEncryptReq 释放AlipayEncryptReq
func ReleaseAlipayEncryptReq(v *AlipayEncryptReq) {
	v.SourceContent = ""
	v.ExternalAppletBizCode = ""
	poolAlipayEncryptReq.Put(v)
}
