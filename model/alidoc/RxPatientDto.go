package alidoc

import (
	"sync"
)

// RxPatientDto 结构体
type RxPatientDto struct {
	// 身份证号
	IdCard string `json:"id_card,omitempty" xml:"id_card,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 年龄
	Age string `json:"age,omitempty" xml:"age,omitempty"`
	// 性别
	Sex string `json:"sex,omitempty" xml:"sex,omitempty"`
}

var poolRxPatientDto = sync.Pool{
	New: func() any {
		return new(RxPatientDto)
	},
}

// GetRxPatientDto() 从对象池中获取RxPatientDto
func GetRxPatientDto() *RxPatientDto {
	return poolRxPatientDto.Get().(*RxPatientDto)
}

// ReleaseRxPatientDto 释放RxPatientDto
func ReleaseRxPatientDto(v *RxPatientDto) {
	v.IdCard = ""
	v.Name = ""
	v.Age = ""
	v.Sex = ""
	poolRxPatientDto.Put(v)
}
