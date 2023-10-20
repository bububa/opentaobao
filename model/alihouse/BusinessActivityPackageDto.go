package alihouse

import (
	"sync"
)

// BusinessActivityPackageDto 结构体
type BusinessActivityPackageDto struct {
	// 签章信息
	SignatureList []BusinessActivitySignatureDto `json:"signature_list,omitempty" xml:"signature_list>business_activity_signature_dto,omitempty"`
	// 协议数据
	AgreementList []BusinessActivityAgreementDto `json:"agreement_list,omitempty" xml:"agreement_list>business_activity_agreement_dto,omitempty"`
	// 结算账户数据
	BankCardList []BusinessActivityBankCardDto `json:"bank_card_list,omitempty" xml:"bank_card_list>business_activity_bank_card_dto,omitempty"`
	// 公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
}

var poolBusinessActivityPackageDto = sync.Pool{
	New: func() any {
		return new(BusinessActivityPackageDto)
	},
}

// GetBusinessActivityPackageDto() 从对象池中获取BusinessActivityPackageDto
func GetBusinessActivityPackageDto() *BusinessActivityPackageDto {
	return poolBusinessActivityPackageDto.Get().(*BusinessActivityPackageDto)
}

// ReleaseBusinessActivityPackageDto 释放BusinessActivityPackageDto
func ReleaseBusinessActivityPackageDto(v *BusinessActivityPackageDto) {
	v.SignatureList = v.SignatureList[:0]
	v.AgreementList = v.AgreementList[:0]
	v.BankCardList = v.BankCardList[:0]
	v.CompanyName = ""
	poolBusinessActivityPackageDto.Put(v)
}
