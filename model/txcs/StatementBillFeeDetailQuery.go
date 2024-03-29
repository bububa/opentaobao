package txcs

import (
	"sync"
)

// StatementBillFeeDetailQuery 结构体
type StatementBillFeeDetailQuery struct {
	// 结算公司编码
	SettlementCompanyCode string `json:"settlement_company_code,omitempty" xml:"settlement_company_code,omitempty"`
	// 账单编码
	StatementBillCode string `json:"statement_bill_code,omitempty" xml:"statement_bill_code,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}

var poolStatementBillFeeDetailQuery = sync.Pool{
	New: func() any {
		return new(StatementBillFeeDetailQuery)
	},
}

// GetStatementBillFeeDetailQuery() 从对象池中获取StatementBillFeeDetailQuery
func GetStatementBillFeeDetailQuery() *StatementBillFeeDetailQuery {
	return poolStatementBillFeeDetailQuery.Get().(*StatementBillFeeDetailQuery)
}

// ReleaseStatementBillFeeDetailQuery 释放StatementBillFeeDetailQuery
func ReleaseStatementBillFeeDetailQuery(v *StatementBillFeeDetailQuery) {
	v.SettlementCompanyCode = ""
	v.StatementBillCode = ""
	v.StartTime = ""
	v.EndTime = ""
	v.PageSize = 0
	v.CurrentPage = 0
	poolStatementBillFeeDetailQuery.Put(v)
}
