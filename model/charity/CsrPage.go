package charity

import (
	"sync"
)

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

var poolCsrPage = sync.Pool{
	New: func() any {
		return new(CsrPage)
	},
}

// GetCsrPage() 从对象池中获取CsrPage
func GetCsrPage() *CsrPage {
	return poolCsrPage.Get().(*CsrPage)
}

// ReleaseCsrPage 释放CsrPage
func ReleaseCsrPage(v *CsrPage) {
	v.List = v.List[:0]
	v.Total = 0
	v.Pages = 0
	v.PageSize = 0
	v.PageNum = 0
	poolCsrPage.Put(v)
}
