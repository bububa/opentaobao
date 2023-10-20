package alihouse

import (
	"sync"
)

// BusinessActivityAgreementDto 结构体
type BusinessActivityAgreementDto struct {
	// 协议内容
	AgreementContent string `json:"agreement_content,omitempty" xml:"agreement_content,omitempty"`
	// 协议名称
	AgreementName string `json:"agreement_name,omitempty" xml:"agreement_name,omitempty"`
	// 协议类型
	AgreementType int64 `json:"agreement_type,omitempty" xml:"agreement_type,omitempty"`
	// ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolBusinessActivityAgreementDto = sync.Pool{
	New: func() any {
		return new(BusinessActivityAgreementDto)
	},
}

// GetBusinessActivityAgreementDto() 从对象池中获取BusinessActivityAgreementDto
func GetBusinessActivityAgreementDto() *BusinessActivityAgreementDto {
	return poolBusinessActivityAgreementDto.Get().(*BusinessActivityAgreementDto)
}

// ReleaseBusinessActivityAgreementDto 释放BusinessActivityAgreementDto
func ReleaseBusinessActivityAgreementDto(v *BusinessActivityAgreementDto) {
	v.AgreementContent = ""
	v.AgreementName = ""
	v.AgreementType = 0
	v.Id = 0
	poolBusinessActivityAgreementDto.Put(v)
}
