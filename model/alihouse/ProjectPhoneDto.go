package alihouse

import (
	"sync"
)

// ProjectPhoneDto 结构体
type ProjectPhoneDto struct {
	// 分机号
	SubPhone string `json:"sub_phone,omitempty" xml:"sub_phone,omitempty"`
	// 主号码
	MainPhone string `json:"main_phone,omitempty" xml:"main_phone,omitempty"`
	// 1-楼盘  2-小区  3-房源  4-VR
	PhoneType string `json:"phone_type,omitempty" xml:"phone_type,omitempty"`
	// 外部楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
}

var poolProjectPhoneDto = sync.Pool{
	New: func() any {
		return new(ProjectPhoneDto)
	},
}

// GetProjectPhoneDto() 从对象池中获取ProjectPhoneDto
func GetProjectPhoneDto() *ProjectPhoneDto {
	return poolProjectPhoneDto.Get().(*ProjectPhoneDto)
}

// ReleaseProjectPhoneDto 释放ProjectPhoneDto
func ReleaseProjectPhoneDto(v *ProjectPhoneDto) {
	v.SubPhone = ""
	v.MainPhone = ""
	v.PhoneType = ""
	v.OuterId = ""
	poolProjectPhoneDto.Put(v)
}
