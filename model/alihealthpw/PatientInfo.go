package alihealthpw

import (
	"sync"
)

// PatientInfo 结构体
type PatientInfo struct {
	// 生日
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 身份证
	PatientCard string `json:"patient_card,omitempty" xml:"patient_card,omitempty"`
	// 性别
	PatientSex string `json:"patient_sex,omitempty" xml:"patient_sex,omitempty"`
	// 姓名
	PatientName string `json:"patient_name,omitempty" xml:"patient_name,omitempty"`
}

var poolPatientInfo = sync.Pool{
	New: func() any {
		return new(PatientInfo)
	},
}

// GetPatientInfo() 从对象池中获取PatientInfo
func GetPatientInfo() *PatientInfo {
	return poolPatientInfo.Get().(*PatientInfo)
}

// ReleasePatientInfo 释放PatientInfo
func ReleasePatientInfo(v *PatientInfo) {
	v.Birthday = ""
	v.PatientCard = ""
	v.PatientSex = ""
	v.PatientName = ""
	poolPatientInfo.Put(v)
}
