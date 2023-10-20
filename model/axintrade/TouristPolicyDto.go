package axintrade

import (
	"sync"
)

// TouristPolicyDto 结构体
type TouristPolicyDto struct {
	// 游客信息填写字段
	FieldTypes []string `json:"field_types,omitempty" xml:"field_types>string,omitempty"`
	// 证件类型
	CertificatesTypes []string `json:"certificates_types,omitempty" xml:"certificates_types>string,omitempty"`
}

var poolTouristPolicyDto = sync.Pool{
	New: func() any {
		return new(TouristPolicyDto)
	},
}

// GetTouristPolicyDto() 从对象池中获取TouristPolicyDto
func GetTouristPolicyDto() *TouristPolicyDto {
	return poolTouristPolicyDto.Get().(*TouristPolicyDto)
}

// ReleaseTouristPolicyDto 释放TouristPolicyDto
func ReleaseTouristPolicyDto(v *TouristPolicyDto) {
	v.FieldTypes = v.FieldTypes[:0]
	v.CertificatesTypes = v.CertificatesTypes[:0]
	poolTouristPolicyDto.Put(v)
}
