package examination

import (
	"sync"
)

// Agreement 结构体
type Agreement struct {
	// 1代表图片2代表文件
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 二级制base64编码1
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}

var poolAgreement = sync.Pool{
	New: func() any {
		return new(Agreement)
	},
}

// GetAgreement() 从对象池中获取Agreement
func GetAgreement() *Agreement {
	return poolAgreement.Get().(*Agreement)
}

// ReleaseAgreement 释放Agreement
func ReleaseAgreement(v *Agreement) {
	v.Type = ""
	v.Content = ""
	poolAgreement.Put(v)
}
