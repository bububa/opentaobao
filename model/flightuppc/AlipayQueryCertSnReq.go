package flightuppc

import (
	"sync"
)

// AlipayQueryCertSnReq 结构体
type AlipayQueryCertSnReq struct {
	// 业务代码
	ExternalAppletBizCode string `json:"external_applet_biz_code,omitempty" xml:"external_applet_biz_code,omitempty"`
	// 证书签名算法
	SignType string `json:"sign_type,omitempty" xml:"sign_type,omitempty"`
}

var poolAlipayQueryCertSnReq = sync.Pool{
	New: func() any {
		return new(AlipayQueryCertSnReq)
	},
}

// GetAlipayQueryCertSnReq() 从对象池中获取AlipayQueryCertSnReq
func GetAlipayQueryCertSnReq() *AlipayQueryCertSnReq {
	return poolAlipayQueryCertSnReq.Get().(*AlipayQueryCertSnReq)
}

// ReleaseAlipayQueryCertSnReq 释放AlipayQueryCertSnReq
func ReleaseAlipayQueryCertSnReq(v *AlipayQueryCertSnReq) {
	v.ExternalAppletBizCode = ""
	v.SignType = ""
	poolAlipayQueryCertSnReq.Put(v)
}
