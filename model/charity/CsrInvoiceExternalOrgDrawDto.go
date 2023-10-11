package charity

// CsrInvoiceExternalOrgDrawDto 结构体
type CsrInvoiceExternalOrgDrawDto struct {
	// 发票文件
	FileList []CsrInvoiceFileDto `json:"file_list,omitempty" xml:"file_list>csr_invoice_file_dto,omitempty"`
	// 发票id
	InvoiceId string `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 开票操作人
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
}
