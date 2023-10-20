package bus

import (
	"sync"
)

// AgentConfirmBookScheduleInfo 结构体
type AgentConfirmBookScheduleInfo struct {
	// 司机姓名
	DriverName string `json:"driver_name,omitempty" xml:"driver_name,omitempty"`
	// 司机联系电话
	DriverPhone string `json:"driver_phone,omitempty" xml:"driver_phone,omitempty"`
	// 车牌号
	LicensePlateNumber string `json:"license_plate_number,omitempty" xml:"license_plate_number,omitempty"`
	// 车型
	MotorcycleType string `json:"motorcycle_type,omitempty" xml:"motorcycle_type,omitempty"`
	// 车辆品牌
	VehicleBrands string `json:"vehicle_brands,omitempty" xml:"vehicle_brands,omitempty"`
}

var poolAgentConfirmBookScheduleInfo = sync.Pool{
	New: func() any {
		return new(AgentConfirmBookScheduleInfo)
	},
}

// GetAgentConfirmBookScheduleInfo() 从对象池中获取AgentConfirmBookScheduleInfo
func GetAgentConfirmBookScheduleInfo() *AgentConfirmBookScheduleInfo {
	return poolAgentConfirmBookScheduleInfo.Get().(*AgentConfirmBookScheduleInfo)
}

// ReleaseAgentConfirmBookScheduleInfo 释放AgentConfirmBookScheduleInfo
func ReleaseAgentConfirmBookScheduleInfo(v *AgentConfirmBookScheduleInfo) {
	v.DriverName = ""
	v.DriverPhone = ""
	v.LicensePlateNumber = ""
	v.MotorcycleType = ""
	v.VehicleBrands = ""
	poolAgentConfirmBookScheduleInfo.Put(v)
}
