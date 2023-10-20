package miniapp

import (
	"sync"
)

// AfterSaleGetWorkTableListResponse 结构体
type AfterSaleGetWorkTableListResponse struct {
	// 调用结果
	Record []AfterSaleGetWorkTableListRecord `json:"record,omitempty" xml:"record>after_sale_get_work_table_list_record,omitempty"`
	// 返回结果数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolAfterSaleGetWorkTableListResponse = sync.Pool{
	New: func() any {
		return new(AfterSaleGetWorkTableListResponse)
	},
}

// GetAfterSaleGetWorkTableListResponse() 从对象池中获取AfterSaleGetWorkTableListResponse
func GetAfterSaleGetWorkTableListResponse() *AfterSaleGetWorkTableListResponse {
	return poolAfterSaleGetWorkTableListResponse.Get().(*AfterSaleGetWorkTableListResponse)
}

// ReleaseAfterSaleGetWorkTableListResponse 释放AfterSaleGetWorkTableListResponse
func ReleaseAfterSaleGetWorkTableListResponse(v *AfterSaleGetWorkTableListResponse) {
	v.Record = v.Record[:0]
	v.TotalCount = 0
	poolAfterSaleGetWorkTableListResponse.Put(v)
}
