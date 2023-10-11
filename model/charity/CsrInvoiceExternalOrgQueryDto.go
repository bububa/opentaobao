package charity

// CsrInvoiceExternalOrgQueryDto 结构体
type CsrInvoiceExternalOrgQueryDto struct {
	// 分页页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
