package txcs

import (
	"sync"
)

// StatementBillConfirmDto 结构体
type StatementBillConfirmDto struct {
	// 账单code
	StatementBillCodes []string `json:"statement_bill_codes,omitempty" xml:"statement_bill_codes>string,omitempty"`
	// 结算公司编码
	SettlementCompanyCode string `json:"settlement_company_code,omitempty" xml:"settlement_company_code,omitempty"`
	// 幂等ID
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作人ID
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 操作人名称
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
}

var poolStatementBillConfirmDto = sync.Pool{
	New: func() any {
		return new(StatementBillConfirmDto)
	},
}

// GetStatementBillConfirmDto() 从对象池中获取StatementBillConfirmDto
func GetStatementBillConfirmDto() *StatementBillConfirmDto {
	return poolStatementBillConfirmDto.Get().(*StatementBillConfirmDto)
}

// ReleaseStatementBillConfirmDto 释放StatementBillConfirmDto
func ReleaseStatementBillConfirmDto(v *StatementBillConfirmDto) {
	v.StatementBillCodes = v.StatementBillCodes[:0]
	v.SettlementCompanyCode = ""
	v.RequestId = ""
	v.OperatorId = ""
	v.OperatorName = ""
	poolStatementBillConfirmDto.Put(v)
}
