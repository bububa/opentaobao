package alidoc

import (
	"sync"
)

// RxDoctorDto 结构体
type RxDoctorDto struct {
	// 医生部门
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 医生姓名
	DoctorName string `json:"doctor_name,omitempty" xml:"doctor_name,omitempty"`
}

var poolRxDoctorDto = sync.Pool{
	New: func() any {
		return new(RxDoctorDto)
	},
}

// GetRxDoctorDto() 从对象池中获取RxDoctorDto
func GetRxDoctorDto() *RxDoctorDto {
	return poolRxDoctorDto.Get().(*RxDoctorDto)
}

// ReleaseRxDoctorDto 释放RxDoctorDto
func ReleaseRxDoctorDto(v *RxDoctorDto) {
	v.DepartName = ""
	v.DoctorName = ""
	poolRxDoctorDto.Put(v)
}
