package consignplatform

import (
	"sync"
)

// PersonDto 结构体
type PersonDto struct {
	// 收件固定电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 收件电话
	MobilePhone string `json:"mobile_phone,omitempty" xml:"mobile_phone,omitempty"`
	// 收件人
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

var poolPersonDto = sync.Pool{
	New: func() any {
		return new(PersonDto)
	},
}

// GetPersonDto() 从对象池中获取PersonDto
func GetPersonDto() *PersonDto {
	return poolPersonDto.Get().(*PersonDto)
}

// ReleasePersonDto 释放PersonDto
func ReleasePersonDto(v *PersonDto) {
	v.Phone = ""
	v.MobilePhone = ""
	v.UserName = ""
	poolPersonDto.Put(v)
}
