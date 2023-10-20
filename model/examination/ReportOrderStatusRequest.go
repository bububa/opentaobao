package examination

import (
	"sync"
)

// ReportOrderStatusRequest 结构体
type ReportOrderStatusRequest struct {
	// 备注
	Note string `json:"note,omitempty" xml:"note,omitempty"`
	// 订单状态，10=待接诊，20=已关闭，30=问诊中，60=已结束
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 外部订单ID
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 医生ID
	DoctorId string `json:"doctor_id,omitempty" xml:"doctor_id,omitempty"`
}

var poolReportOrderStatusRequest = sync.Pool{
	New: func() any {
		return new(ReportOrderStatusRequest)
	},
}

// GetReportOrderStatusRequest() 从对象池中获取ReportOrderStatusRequest
func GetReportOrderStatusRequest() *ReportOrderStatusRequest {
	return poolReportOrderStatusRequest.Get().(*ReportOrderStatusRequest)
}

// ReleaseReportOrderStatusRequest 释放ReportOrderStatusRequest
func ReleaseReportOrderStatusRequest(v *ReportOrderStatusRequest) {
	v.Note = ""
	v.Status = ""
	v.OrderId = ""
	v.OuterOrderId = ""
	v.DoctorId = ""
	poolReportOrderStatusRequest.Put(v)
}
