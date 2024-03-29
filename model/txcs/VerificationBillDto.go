package txcs

import (
	"sync"
)

// VerificationBillDto 结构体
type VerificationBillDto struct {
	// 发票信息列表
	InvoiceInfoDTOs []InvoiceInfoDto `json:"invoice_info_d_t_os,omitempty" xml:"invoice_info_d_t_os>invoice_info_dto,omitempty"`
	// 对账单编号列表
	StatementBillCodes []string `json:"statement_bill_codes,omitempty" xml:"statement_bill_codes>string,omitempty"`
	// 结算公司
	SettlementCompanyCode string `json:"settlement_company_code,omitempty" xml:"settlement_company_code,omitempty"`
	// 备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 操作人名称
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 幂等ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 核销单类型
	VerificationType string `json:"verification_type,omitempty" xml:"verification_type,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 操作人ID
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
}

var poolVerificationBillDto = sync.Pool{
	New: func() any {
		return new(VerificationBillDto)
	},
}

// GetVerificationBillDto() 从对象池中获取VerificationBillDto
func GetVerificationBillDto() *VerificationBillDto {
	return poolVerificationBillDto.Get().(*VerificationBillDto)
}

// ReleaseVerificationBillDto 释放VerificationBillDto
func ReleaseVerificationBillDto(v *VerificationBillDto) {
	v.InvoiceInfoDTOs = v.InvoiceInfoDTOs[:0]
	v.StatementBillCodes = v.StatementBillCodes[:0]
	v.SettlementCompanyCode = ""
	v.Memo = ""
	v.OperatorName = ""
	v.RequestId = ""
	v.VerificationType = ""
	v.Currency = ""
	v.OperatorId = ""
	poolVerificationBillDto.Put(v)
}
