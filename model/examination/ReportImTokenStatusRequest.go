package examination

import (
	"sync"
)

// ReportImTokenStatusRequest 结构体
type ReportImTokenStatusRequest struct {
	// 令牌值
	ImToken string `json:"im_token,omitempty" xml:"im_token,omitempty"`
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolReportImTokenStatusRequest = sync.Pool{
	New: func() any {
		return new(ReportImTokenStatusRequest)
	},
}

// GetReportImTokenStatusRequest() 从对象池中获取ReportImTokenStatusRequest
func GetReportImTokenStatusRequest() *ReportImTokenStatusRequest {
	return poolReportImTokenStatusRequest.Get().(*ReportImTokenStatusRequest)
}

// ReleaseReportImTokenStatusRequest 释放ReportImTokenStatusRequest
func ReleaseReportImTokenStatusRequest(v *ReportImTokenStatusRequest) {
	v.ImToken = ""
	v.OrderId = ""
	poolReportImTokenStatusRequest.Put(v)
}
