package drug

import (
	"sync"
)

// MedicalInsurancePaymentDto 结构体
type MedicalInsurancePaymentDto struct {
	// 医保支付金额信息
	FundBillList []FundBillDto `json:"fund_bill_list,omitempty" xml:"fund_bill_list>fund_bill_dto,omitempty"`
	// 医保局返回的医保支付内容
	IndustrySepcDetail string `json:"industry_sepc_detail,omitempty" xml:"industry_sepc_detail,omitempty"`
}

var poolMedicalInsurancePaymentDto = sync.Pool{
	New: func() any {
		return new(MedicalInsurancePaymentDto)
	},
}

// GetMedicalInsurancePaymentDto() 从对象池中获取MedicalInsurancePaymentDto
func GetMedicalInsurancePaymentDto() *MedicalInsurancePaymentDto {
	return poolMedicalInsurancePaymentDto.Get().(*MedicalInsurancePaymentDto)
}

// ReleaseMedicalInsurancePaymentDto 释放MedicalInsurancePaymentDto
func ReleaseMedicalInsurancePaymentDto(v *MedicalInsurancePaymentDto) {
	v.FundBillList = v.FundBillList[:0]
	v.IndustrySepcDetail = ""
	poolMedicalInsurancePaymentDto.Put(v)
}
