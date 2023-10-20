package flightuppc

import (
	"sync"
)

// AlipayCheckSignReq 结构体
type AlipayCheckSignReq struct {
	// 加签自定义参数，格式遵循http请求路径参数的格式
	SourceContent string `json:"source_content,omitempty" xml:"source_content,omitempty"`
	// 签名
	Signature string `json:"signature,omitempty" xml:"signature,omitempty"`
	// 业务代码
	ExternalAppletBizCode string `json:"external_applet_biz_code,omitempty" xml:"external_applet_biz_code,omitempty"`
	// 签名算法类型
	SignType string `json:"sign_type,omitempty" xml:"sign_type,omitempty"`
}

var poolAlipayCheckSignReq = sync.Pool{
	New: func() any {
		return new(AlipayCheckSignReq)
	},
}

// GetAlipayCheckSignReq() 从对象池中获取AlipayCheckSignReq
func GetAlipayCheckSignReq() *AlipayCheckSignReq {
	return poolAlipayCheckSignReq.Get().(*AlipayCheckSignReq)
}

// ReleaseAlipayCheckSignReq 释放AlipayCheckSignReq
func ReleaseAlipayCheckSignReq(v *AlipayCheckSignReq) {
	v.SourceContent = ""
	v.Signature = ""
	v.ExternalAppletBizCode = ""
	v.SignType = ""
	poolAlipayCheckSignReq.Put(v)
}
