package examination

import (
	"sync"
)

// IsvGetReportDiagnoseUrlRequest 结构体
type IsvGetReportDiagnoseUrlRequest struct {
	// 服务订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolIsvGetReportDiagnoseUrlRequest = sync.Pool{
	New: func() any {
		return new(IsvGetReportDiagnoseUrlRequest)
	},
}

// GetIsvGetReportDiagnoseUrlRequest() 从对象池中获取IsvGetReportDiagnoseUrlRequest
func GetIsvGetReportDiagnoseUrlRequest() *IsvGetReportDiagnoseUrlRequest {
	return poolIsvGetReportDiagnoseUrlRequest.Get().(*IsvGetReportDiagnoseUrlRequest)
}

// ReleaseIsvGetReportDiagnoseUrlRequest 释放IsvGetReportDiagnoseUrlRequest
func ReleaseIsvGetReportDiagnoseUrlRequest(v *IsvGetReportDiagnoseUrlRequest) {
	v.OrderId = ""
	poolIsvGetReportDiagnoseUrlRequest.Put(v)
}
