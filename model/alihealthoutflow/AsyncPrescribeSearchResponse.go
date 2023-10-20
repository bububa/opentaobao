package alihealthoutflow

import (
	"sync"
)

// AsyncPrescribeSearchResponse 结构体
type AsyncPrescribeSearchResponse struct {
	// 处方号列表
	RxNos []string `json:"rx_nos,omitempty" xml:"rx_nos>string,omitempty"`
	// 第几页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页多少条
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总数量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 开始时间
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

var poolAsyncPrescribeSearchResponse = sync.Pool{
	New: func() any {
		return new(AsyncPrescribeSearchResponse)
	},
}

// GetAsyncPrescribeSearchResponse() 从对象池中获取AsyncPrescribeSearchResponse
func GetAsyncPrescribeSearchResponse() *AsyncPrescribeSearchResponse {
	return poolAsyncPrescribeSearchResponse.Get().(*AsyncPrescribeSearchResponse)
}

// ReleaseAsyncPrescribeSearchResponse 释放AsyncPrescribeSearchResponse
func ReleaseAsyncPrescribeSearchResponse(v *AsyncPrescribeSearchResponse) {
	v.RxNos = v.RxNos[:0]
	v.PageNo = 0
	v.PageSize = 0
	v.Total = 0
	v.StartTime = 0
	v.EndTime = 0
	poolAsyncPrescribeSearchResponse.Put(v)
}
