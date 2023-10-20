package examination

import (
	"sync"
)

// ReportImTokenStatusResponse 结构体
type ReportImTokenStatusResponse struct {
	// 令牌校验失败原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 令牌校验状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolReportImTokenStatusResponse = sync.Pool{
	New: func() any {
		return new(ReportImTokenStatusResponse)
	},
}

// GetReportImTokenStatusResponse() 从对象池中获取ReportImTokenStatusResponse
func GetReportImTokenStatusResponse() *ReportImTokenStatusResponse {
	return poolReportImTokenStatusResponse.Get().(*ReportImTokenStatusResponse)
}

// ReleaseReportImTokenStatusResponse 释放ReportImTokenStatusResponse
func ReleaseReportImTokenStatusResponse(v *ReportImTokenStatusResponse) {
	v.Message = ""
	v.Status = ""
	poolReportImTokenStatusResponse.Put(v)
}
