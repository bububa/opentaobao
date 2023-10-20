package miniapp

import (
	"sync"
)

// AfterSaleTableSelectResponse 结构体
type AfterSaleTableSelectResponse struct {
	// 查询结果
	JsonRecord string `json:"json_record,omitempty" xml:"json_record,omitempty"`
	// 查询结果条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolAfterSaleTableSelectResponse = sync.Pool{
	New: func() any {
		return new(AfterSaleTableSelectResponse)
	},
}

// GetAfterSaleTableSelectResponse() 从对象池中获取AfterSaleTableSelectResponse
func GetAfterSaleTableSelectResponse() *AfterSaleTableSelectResponse {
	return poolAfterSaleTableSelectResponse.Get().(*AfterSaleTableSelectResponse)
}

// ReleaseAfterSaleTableSelectResponse 释放AfterSaleTableSelectResponse
func ReleaseAfterSaleTableSelectResponse(v *AfterSaleTableSelectResponse) {
	v.JsonRecord = ""
	v.TotalCount = 0
	poolAfterSaleTableSelectResponse.Put(v)
}
