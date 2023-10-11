package charity

// CsrPage 结构体
type CsrPage struct {
	// 分页数据
	List []CsrInvoiceExternalOrgQueryResultDto `json:"list,omitempty" xml:"list>csr_invoice_external_org_query_result_dto,omitempty"`
	// 总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 总页数
	Pages int64 `json:"pages,omitempty" xml:"pages,omitempty"`
	// 分页大小，默认20，最大100
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页码
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
