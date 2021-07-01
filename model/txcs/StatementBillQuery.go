package txcs

// StatementBillQuery 结构体
type StatementBillQuery struct {
	// 对账单号
	StatementBillCode string `json:"statement_bill_code,omitempty" xml:"statement_bill_code,omitempty"`
}
