package miniapp

import (
	"sync"
)

// AfterSaleFieldMetaResponse 结构体
type AfterSaleFieldMetaResponse struct {
	//
	Record []AfterSaleFieldMetaRecord `json:"record,omitempty" xml:"record>after_sale_field_meta_record,omitempty"`
	// 返回结果条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolAfterSaleFieldMetaResponse = sync.Pool{
	New: func() any {
		return new(AfterSaleFieldMetaResponse)
	},
}

// GetAfterSaleFieldMetaResponse() 从对象池中获取AfterSaleFieldMetaResponse
func GetAfterSaleFieldMetaResponse() *AfterSaleFieldMetaResponse {
	return poolAfterSaleFieldMetaResponse.Get().(*AfterSaleFieldMetaResponse)
}

// ReleaseAfterSaleFieldMetaResponse 释放AfterSaleFieldMetaResponse
func ReleaseAfterSaleFieldMetaResponse(v *AfterSaleFieldMetaResponse) {
	v.Record = v.Record[:0]
	v.TotalCount = 0
	poolAfterSaleFieldMetaResponse.Put(v)
}
