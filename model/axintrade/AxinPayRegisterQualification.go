package axintrade

import (
	"sync"
)

// AxinPayRegisterQualification 结构体
type AxinPayRegisterQualification struct {
	// 行业资格证图片ID
	ImageId string `json:"image_id,omitempty" xml:"image_id,omitempty"`
	// 类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
}

var poolAxinPayRegisterQualification = sync.Pool{
	New: func() any {
		return new(AxinPayRegisterQualification)
	},
}

// GetAxinPayRegisterQualification() 从对象池中获取AxinPayRegisterQualification
func GetAxinPayRegisterQualification() *AxinPayRegisterQualification {
	return poolAxinPayRegisterQualification.Get().(*AxinPayRegisterQualification)
}

// ReleaseAxinPayRegisterQualification 释放AxinPayRegisterQualification
func ReleaseAxinPayRegisterQualification(v *AxinPayRegisterQualification) {
	v.ImageId = ""
	v.Type = ""
	poolAxinPayRegisterQualification.Put(v)
}
