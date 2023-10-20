package alihealthpw

import (
	"sync"
)

// TreatInfo 结构体
type TreatInfo struct {
	// 诊疗报告
	DiagnosisReports []string `json:"diagnosis_reports,omitempty" xml:"diagnosis_reports>string,omitempty"`
	// 就诊医生
	Doctor string `json:"doctor,omitempty" xml:"doctor,omitempty"`
	// 医院
	Hospital string `json:"hospital,omitempty" xml:"hospital,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 疾病名
	DiseaseName string `json:"disease_name,omitempty" xml:"disease_name,omitempty"`
	// 药品名称
	DrugName string `json:"drug_name,omitempty" xml:"drug_name,omitempty"`
}

var poolTreatInfo = sync.Pool{
	New: func() any {
		return new(TreatInfo)
	},
}

// GetTreatInfo() 从对象池中获取TreatInfo
func GetTreatInfo() *TreatInfo {
	return poolTreatInfo.Get().(*TreatInfo)
}

// ReleaseTreatInfo 释放TreatInfo
func ReleaseTreatInfo(v *TreatInfo) {
	v.DiagnosisReports = v.DiagnosisReports[:0]
	v.Doctor = ""
	v.Hospital = ""
	v.City = ""
	v.DiseaseName = ""
	v.DrugName = ""
	poolTreatInfo.Put(v)
}
