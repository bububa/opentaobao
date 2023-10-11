package wlbimports

// SelfMailTruckInfoRequest 结构体
type SelfMailTruckInfoRequest struct {
	// 司机姓名
	DriverName string `json:"driver_name,omitempty" xml:"driver_name,omitempty"`
	// 司机手机号
	DriverMobilePhone string `json:"driver_mobile_phone,omitempty" xml:"driver_mobile_phone,omitempty"`
	// 车牌（自寄卡车模式必填）
	TruckNum string `json:"truck_num,omitempty" xml:"truck_num,omitempty"`
	// 预计送达时间（自寄卡车模式必填）
	ExpectedArriveTime int64 `json:"expected_arrive_time,omitempty" xml:"expected_arrive_time,omitempty"`
	// 预计揽收时间（自寄卡车模式必填）
	ExpectedPickupTime int64 `json:"expected_pickup_time,omitempty" xml:"expected_pickup_time,omitempty"`
}
