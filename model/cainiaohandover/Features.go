package cainiaohandover

import (
	"sync"
)

// Features 结构体
type Features struct {
	// 容器类型(1、托盘;2、大包或盒子3、散装)
	ContainerType string `json:"container_type,omitempty" xml:"container_type,omitempty"`
	// 是否预先组大包，true：是。false：否
	PrePackage string `json:"pre_package,omitempty" xml:"pre_package,omitempty"`
	// 自送仓资源编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 自送仓资源所对应的schemeCode
	ServiceResourceCode string `json:"service_resource_code,omitempty" xml:"service_resource_code,omitempty"`
	// 自寄时快递公司运单号
	ExpressMailNo string `json:"express_mail_no,omitempty" xml:"express_mail_no,omitempty"`
	// 自寄时快递公司的ID
	ExpressCompanyId string `json:"express_company_id,omitempty" xml:"express_company_id,omitempty"`
	// 自寄时快递公司的名称
	ExpressCompanyName string `json:"express_company_name,omitempty" xml:"express_company_name,omitempty"`
	// 预约流程
	AppointmentProcess string `json:"appointment_process,omitempty" xml:"appointment_process,omitempty"`
	// 揽收时间
	PickupWorkTime string `json:"pickup_work_time,omitempty" xml:"pickup_work_time,omitempty"`
	// 货好时间时间戳
	GmtReadyToShip int64 `json:"gmt_ready_to_ship,omitempty" xml:"gmt_ready_to_ship,omitempty"`
	// 托盘数量
	PalletQuantity int64 `json:"pallet_quantity,omitempty" xml:"pallet_quantity,omitempty"`
	// 是否需要预约
	NeedAppointment bool `json:"need_appointment,omitempty" xml:"need_appointment,omitempty"`
}

var poolFeatures = sync.Pool{
	New: func() any {
		return new(Features)
	},
}

// GetFeatures() 从对象池中获取Features
func GetFeatures() *Features {
	return poolFeatures.Get().(*Features)
}

// ReleaseFeatures 释放Features
func ReleaseFeatures(v *Features) {
	v.ContainerType = ""
	v.PrePackage = ""
	v.WarehouseCode = ""
	v.ServiceResourceCode = ""
	v.ExpressMailNo = ""
	v.ExpressCompanyId = ""
	v.ExpressCompanyName = ""
	v.AppointmentProcess = ""
	v.PickupWorkTime = ""
	v.GmtReadyToShip = 0
	v.PalletQuantity = 0
	v.NeedAppointment = false
	poolFeatures.Put(v)
}
