package txcs

import (
	"sync"
)

// StatementBillQuery 结构体
type StatementBillQuery struct {
	// 对账单号
	StatementBillCode string `json:"statement_bill_code,omitempty" xml:"statement_bill_code,omitempty"`
}

var poolStatementBillQuery = sync.Pool{
	New: func() any {
		return new(StatementBillQuery)
	},
}

// GetStatementBillQuery() 从对象池中获取StatementBillQuery
func GetStatementBillQuery() *StatementBillQuery {
	return poolStatementBillQuery.Get().(*StatementBillQuery)
}

// ReleaseStatementBillQuery 释放StatementBillQuery
func ReleaseStatementBillQuery(v *StatementBillQuery) {
	v.StatementBillCode = ""
	poolStatementBillQuery.Put(v)
}
