package alidoc

import (
	"sync"
)

// DiseaseInfo 结构体
type DiseaseInfo struct {
	// 诊断名称
	DiseaseName string `json:"disease_name,omitempty" xml:"disease_name,omitempty"`
	// 诊断code
	DiseaseCode string `json:"disease_code,omitempty" xml:"disease_code,omitempty"`
}

var poolDiseaseInfo = sync.Pool{
	New: func() any {
		return new(DiseaseInfo)
	},
}

// GetDiseaseInfo() 从对象池中获取DiseaseInfo
func GetDiseaseInfo() *DiseaseInfo {
	return poolDiseaseInfo.Get().(*DiseaseInfo)
}

// ReleaseDiseaseInfo 释放DiseaseInfo
func ReleaseDiseaseInfo(v *DiseaseInfo) {
	v.DiseaseName = ""
	v.DiseaseCode = ""
	poolDiseaseInfo.Put(v)
}
