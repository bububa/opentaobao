package vaccin

import (
	"sync"
)

// VaccineInfo 结构体
type VaccineInfo struct {
	// 疫苗名称
	VaccineName string `json:"vaccine_name,omitempty" xml:"vaccine_name,omitempty"`
	// 疫苗编码
	VaccineCode string `json:"vaccine_code,omitempty" xml:"vaccine_code,omitempty"`
	// 阿里健康疫苗编码
	VaccineGbCode string `json:"vaccine_gb_code,omitempty" xml:"vaccine_gb_code,omitempty"`
	// 疫苗针次
	TheTimes int64 `json:"the_times,omitempty" xml:"the_times,omitempty"`
}

var poolVaccineInfo = sync.Pool{
	New: func() any {
		return new(VaccineInfo)
	},
}

// GetVaccineInfo() 从对象池中获取VaccineInfo
func GetVaccineInfo() *VaccineInfo {
	return poolVaccineInfo.Get().(*VaccineInfo)
}

// ReleaseVaccineInfo 释放VaccineInfo
func ReleaseVaccineInfo(v *VaccineInfo) {
	v.VaccineName = ""
	v.VaccineCode = ""
	v.VaccineGbCode = ""
	v.TheTimes = 0
	poolVaccineInfo.Put(v)
}
