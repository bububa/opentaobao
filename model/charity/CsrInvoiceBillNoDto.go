package charity

// CsrInvoiceBillNoDto 结构体
type CsrInvoiceBillNoDto struct {
	// 账单编号
	BillId string `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
	// 账单时间
	BillTime int64 `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
}
