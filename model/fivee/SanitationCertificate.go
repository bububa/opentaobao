package fivee

import (
	"sync"
)

// SanitationCertificate 结构体
type SanitationCertificate struct {
	// 批次信息
	BatchProducts []BatchProduct `json:"batch_products,omitempty" xml:"batch_products>batch_product,omitempty"`
	// 下载地址列表
	Urls []string `json:"urls,omitempty" xml:"urls>string,omitempty"`
	// 编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolSanitationCertificate = sync.Pool{
	New: func() any {
		return new(SanitationCertificate)
	},
}

// GetSanitationCertificate() 从对象池中获取SanitationCertificate
func GetSanitationCertificate() *SanitationCertificate {
	return poolSanitationCertificate.Get().(*SanitationCertificate)
}

// ReleaseSanitationCertificate 释放SanitationCertificate
func ReleaseSanitationCertificate(v *SanitationCertificate) {
	v.BatchProducts = v.BatchProducts[:0]
	v.Urls = v.Urls[:0]
	v.Code = ""
	poolSanitationCertificate.Put(v)
}
