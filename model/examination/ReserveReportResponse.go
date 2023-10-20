package examination

import (
	"sync"
)

// ReserveReportResponse 结构体
type ReserveReportResponse struct {
	// 11
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 11
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
}

var poolReserveReportResponse = sync.Pool{
	New: func() any {
		return new(ReserveReportResponse)
	},
}

// GetReserveReportResponse() 从对象池中获取ReserveReportResponse
func GetReserveReportResponse() *ReserveReportResponse {
	return poolReserveReportResponse.Get().(*ReserveReportResponse)
}

// ReleaseReserveReportResponse 释放ReserveReportResponse
func ReleaseReserveReportResponse(v *ReserveReportResponse) {
	v.Msg = ""
	v.ResponseCode = ""
	poolReserveReportResponse.Put(v)
}
