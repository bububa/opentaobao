package fivee

import (
	"sync"
)

// Licence 结构体
type Licence struct {
	// 附件下载地址列表
	Urls []string `json:"urls,omitempty" xml:"urls>string,omitempty"`
	// 认证机构
	CertificationBody string `json:"certification_body,omitempty" xml:"certification_body,omitempty"`
	// 编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 到期日期
	DueDate string `json:"due_date,omitempty" xml:"due_date,omitempty"`
	// 生效日期
	EffectiveDate string `json:"effective_date,omitempty" xml:"effective_date,omitempty"`
	// 证照名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolLicence = sync.Pool{
	New: func() any {
		return new(Licence)
	},
}

// GetLicence() 从对象池中获取Licence
func GetLicence() *Licence {
	return poolLicence.Get().(*Licence)
}

// ReleaseLicence 释放Licence
func ReleaseLicence(v *Licence) {
	v.Urls = v.Urls[:0]
	v.CertificationBody = ""
	v.Code = ""
	v.DueDate = ""
	v.EffectiveDate = ""
	v.Name = ""
	v.Type = 0
	poolLicence.Put(v)
}
