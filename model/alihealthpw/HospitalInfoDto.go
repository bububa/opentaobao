package alihealthpw

import (
	"sync"
)

// HospitalInfoDto 结构体
type HospitalInfoDto struct {
	// 医院联系方式
	HospitalPhone string `json:"hospital_phone,omitempty" xml:"hospital_phone,omitempty"`
	// 医院名称
	HospitalName string `json:"hospital_name,omitempty" xml:"hospital_name,omitempty"`
	// 医院地址
	HospitalAddress string `json:"hospital_address,omitempty" xml:"hospital_address,omitempty"`
	// 医院logo
	HospitalLogo string `json:"hospital_logo,omitempty" xml:"hospital_logo,omitempty"`
	// 医院等级
	HospitalGrade string `json:"hospital_grade,omitempty" xml:"hospital_grade,omitempty"`
	// 医院编码
	HospitalCode string `json:"hospital_code,omitempty" xml:"hospital_code,omitempty"`
}

var poolHospitalInfoDto = sync.Pool{
	New: func() any {
		return new(HospitalInfoDto)
	},
}

// GetHospitalInfoDto() 从对象池中获取HospitalInfoDto
func GetHospitalInfoDto() *HospitalInfoDto {
	return poolHospitalInfoDto.Get().(*HospitalInfoDto)
}

// ReleaseHospitalInfoDto 释放HospitalInfoDto
func ReleaseHospitalInfoDto(v *HospitalInfoDto) {
	v.HospitalPhone = ""
	v.HospitalName = ""
	v.HospitalAddress = ""
	v.HospitalLogo = ""
	v.HospitalGrade = ""
	v.HospitalCode = ""
	poolHospitalInfoDto.Put(v)
}
