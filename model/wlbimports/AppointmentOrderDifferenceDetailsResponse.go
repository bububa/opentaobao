package wlbimports

import (
	"sync"
)

// AppointmentOrderDifferenceDetailsResponse 结构体
type AppointmentOrderDifferenceDetailsResponse struct {
	// 预约单号
	ParcelList []Parcel `json:"parcel_list,omitempty" xml:"parcel_list>parcel,omitempty"`
	// 预约单号HOxxxxxxxxx
	HandoverOrderCode string `json:"handover_order_code,omitempty" xml:"handover_order_code,omitempty"`
	// 同预约单号
	TrackingNumber string `json:"tracking_number,omitempty" xml:"tracking_number,omitempty"`
	// 0-全部小包接收成功
	OpCode string `json:"op_code,omitempty" xml:"op_code,omitempty"`
	// he operator CP&#39;s resource code
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// Operator contac（name or phone or email
	OperatorContact string `json:"operator_contact,omitempty" xml:"operator_contact,omitempty"`
	// 备注
	OpRemark string `json:"op_remark,omitempty" xml:"op_remark,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 符合条件的总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前页
	CurrentPageIndex int64 `json:"current_page_index,omitempty" xml:"current_page_index,omitempty"`
}

var poolAppointmentOrderDifferenceDetailsResponse = sync.Pool{
	New: func() any {
		return new(AppointmentOrderDifferenceDetailsResponse)
	},
}

// GetAppointmentOrderDifferenceDetailsResponse() 从对象池中获取AppointmentOrderDifferenceDetailsResponse
func GetAppointmentOrderDifferenceDetailsResponse() *AppointmentOrderDifferenceDetailsResponse {
	return poolAppointmentOrderDifferenceDetailsResponse.Get().(*AppointmentOrderDifferenceDetailsResponse)
}

// ReleaseAppointmentOrderDifferenceDetailsResponse 释放AppointmentOrderDifferenceDetailsResponse
func ReleaseAppointmentOrderDifferenceDetailsResponse(v *AppointmentOrderDifferenceDetailsResponse) {
	v.ParcelList = v.ParcelList[:0]
	v.HandoverOrderCode = ""
	v.TrackingNumber = ""
	v.OpCode = ""
	v.Operator = ""
	v.OperatorContact = ""
	v.OpRemark = ""
	v.PageSize = 0
	v.TotalCount = 0
	v.CurrentPageIndex = 0
	poolAppointmentOrderDifferenceDetailsResponse.Put(v)
}
