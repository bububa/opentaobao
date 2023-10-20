package paimai

import (
	"sync"
)

// VehicleDetectServerReport4Top 结构体
type VehicleDetectServerReport4Top struct {
	// 拍卖服务单单号（与检测单单号不能同时为空）
	ServiceCaseNo string `json:"service_case_no,omitempty" xml:"service_case_no,omitempty"`
	// 检测机构检测单单号（与拍卖服务单单号，不能同时为空）
	OrderNum string `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

var poolVehicleDetectServerReport4Top = sync.Pool{
	New: func() any {
		return new(VehicleDetectServerReport4Top)
	},
}

// GetVehicleDetectServerReport4Top() 从对象池中获取VehicleDetectServerReport4Top
func GetVehicleDetectServerReport4Top() *VehicleDetectServerReport4Top {
	return poolVehicleDetectServerReport4Top.Get().(*VehicleDetectServerReport4Top)
}

// ReleaseVehicleDetectServerReport4Top 释放VehicleDetectServerReport4Top
func ReleaseVehicleDetectServerReport4Top(v *VehicleDetectServerReport4Top) {
	v.ServiceCaseNo = ""
	v.OrderNum = ""
	poolVehicleDetectServerReport4Top.Put(v)
}
