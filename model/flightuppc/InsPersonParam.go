package flightuppc

import (
	"sync"
)

// InsPersonParam 结构体
type InsPersonParam struct {
	// 证件号码
	CertNo string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 身份类型
	IdentityType string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
	// 证件名称
	CertName string `json:"cert_name,omitempty" xml:"cert_name,omitempty"`
	// 生日，不需要传
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 证件类型，不需要传
	CertType string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// 电话号码，不需要传
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
}

var poolInsPersonParam = sync.Pool{
	New: func() any {
		return new(InsPersonParam)
	},
}

// GetInsPersonParam() 从对象池中获取InsPersonParam
func GetInsPersonParam() *InsPersonParam {
	return poolInsPersonParam.Get().(*InsPersonParam)
}

// ReleaseInsPersonParam 释放InsPersonParam
func ReleaseInsPersonParam(v *InsPersonParam) {
	v.CertNo = ""
	v.IdentityType = ""
	v.CertName = ""
	v.Birthday = ""
	v.CertType = ""
	v.Phone = ""
	poolInsPersonParam.Put(v)
}
