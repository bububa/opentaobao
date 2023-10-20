package ott

import (
	"sync"
)

// CertificationDto 结构体
type CertificationDto struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolCertificationDto = sync.Pool{
	New: func() any {
		return new(CertificationDto)
	},
}

// GetCertificationDto() 从对象池中获取CertificationDto
func GetCertificationDto() *CertificationDto {
	return poolCertificationDto.Get().(*CertificationDto)
}

// ReleaseCertificationDto 释放CertificationDto
func ReleaseCertificationDto(v *CertificationDto) {
	v.Code = ""
	v.Value = ""
	poolCertificationDto.Put(v)
}
