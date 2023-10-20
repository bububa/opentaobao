package flight

import (
	"sync"
)

// PassengerDto 结构体
type PassengerDto struct {
	// 证件类型
	DocumentsType []string `json:"documents_type,omitempty" xml:"documents_type>string,omitempty"`
	// 年龄限制
	AgeLimit string `json:"age_limit,omitempty" xml:"age_limit,omitempty"`
	// 身份地域限制
	DocumentsLimit string `json:"documents_limit,omitempty" xml:"documents_limit,omitempty"`
	// 任务限制
	PaxNum string `json:"pax_num,omitempty" xml:"pax_num,omitempty"`
	// 产品类型
	ProductCode int64 `json:"product_code,omitempty" xml:"product_code,omitempty"`
}

var poolPassengerDto = sync.Pool{
	New: func() any {
		return new(PassengerDto)
	},
}

// GetPassengerDto() 从对象池中获取PassengerDto
func GetPassengerDto() *PassengerDto {
	return poolPassengerDto.Get().(*PassengerDto)
}

// ReleasePassengerDto 释放PassengerDto
func ReleasePassengerDto(v *PassengerDto) {
	v.DocumentsType = v.DocumentsType[:0]
	v.AgeLimit = ""
	v.DocumentsLimit = ""
	v.PaxNum = ""
	v.ProductCode = 0
	poolPassengerDto.Put(v)
}
