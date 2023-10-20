package paimai

import (
	"sync"
)

// VehicleServerOrderInfo4Top 结构体
type VehicleServerOrderInfo4Top struct {
	// 拍卖服务单单号（与检测单单号不能同时为空）
	ServiceCaseNo string `json:"service_case_no,omitempty" xml:"service_case_no,omitempty"`
	// 检测机构检测单单号（与拍卖服务单单号，不能同时为空）
	OrderNum string `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// 取消原因（取消时不允许为空）
	CancelReason string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
	// 同步状态（-1：取消订单 0：同步检测师傅信息 1：检测中 2：检测完成）
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 取消类型，取消时不允许为空（ 0：取消并且全额退款； 1：取消 不退还所付金额）
	CancelType int64 `json:"cancel_type,omitempty" xml:"cancel_type,omitempty"`
}

var poolVehicleServerOrderInfo4Top = sync.Pool{
	New: func() any {
		return new(VehicleServerOrderInfo4Top)
	},
}

// GetVehicleServerOrderInfo4Top() 从对象池中获取VehicleServerOrderInfo4Top
func GetVehicleServerOrderInfo4Top() *VehicleServerOrderInfo4Top {
	return poolVehicleServerOrderInfo4Top.Get().(*VehicleServerOrderInfo4Top)
}

// ReleaseVehicleServerOrderInfo4Top 释放VehicleServerOrderInfo4Top
func ReleaseVehicleServerOrderInfo4Top(v *VehicleServerOrderInfo4Top) {
	v.ServiceCaseNo = ""
	v.OrderNum = ""
	v.CancelReason = ""
	v.Status = 0
	v.CancelType = 0
	poolVehicleServerOrderInfo4Top.Put(v)
}
