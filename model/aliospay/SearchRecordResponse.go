package aliospay

import (
	"sync"
)

// SearchRecordResponse 结构体
type SearchRecordResponse struct {
	// 支付记录列表
	Datas []PayRecordData `json:"datas,omitempty" xml:"datas>pay_record_data,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolSearchRecordResponse = sync.Pool{
	New: func() any {
		return new(SearchRecordResponse)
	},
}

// GetSearchRecordResponse() 从对象池中获取SearchRecordResponse
func GetSearchRecordResponse() *SearchRecordResponse {
	return poolSearchRecordResponse.Get().(*SearchRecordResponse)
}

// ReleaseSearchRecordResponse 释放SearchRecordResponse
func ReleaseSearchRecordResponse(v *SearchRecordResponse) {
	v.Datas = v.Datas[:0]
	v.Total = 0
	poolSearchRecordResponse.Put(v)
}
