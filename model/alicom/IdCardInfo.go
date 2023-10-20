package alicom

import (
	"sync"
)

// IdCardInfo 结构体
type IdCardInfo struct {
	// 证件号码
	CardNumber string `json:"card_number,omitempty" xml:"card_number,omitempty"`
	// 证件类型
	CardType string `json:"card_type,omitempty" xml:"card_type,omitempty"`
	// 证件名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 身份证反面照
	BackCertPic string `json:"back_cert_pic,omitempty" xml:"back_cert_pic,omitempty"`
	// 身份证正面照
	FaceCertPic string `json:"face_cert_pic,omitempty" xml:"face_cert_pic,omitempty"`
	// 手持证件照
	HoldCertPic string `json:"hold_cert_pic,omitempty" xml:"hold_cert_pic,omitempty"`
}

var poolIdCardInfo = sync.Pool{
	New: func() any {
		return new(IdCardInfo)
	},
}

// GetIdCardInfo() 从对象池中获取IdCardInfo
func GetIdCardInfo() *IdCardInfo {
	return poolIdCardInfo.Get().(*IdCardInfo)
}

// ReleaseIdCardInfo 释放IdCardInfo
func ReleaseIdCardInfo(v *IdCardInfo) {
	v.CardNumber = ""
	v.CardType = ""
	v.Name = ""
	v.BackCertPic = ""
	v.FaceCertPic = ""
	v.HoldCertPic = ""
	poolIdCardInfo.Put(v)
}
