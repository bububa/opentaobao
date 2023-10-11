package charity

// CsrInvoiceExternalOrgRejectDto 结构体
type CsrInvoiceExternalOrgRejectDto struct {
	// 拒绝原因
	RejectReason string `json:"reject_reason,omitempty" xml:"reject_reason,omitempty"`
	// 发票id
	InvoiceId string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
}
