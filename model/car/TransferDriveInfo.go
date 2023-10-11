package car

// TransferDriveInfo 结构体
type TransferDriveInfo struct {
	// 车牌号
	License string `json:"license,omitempty" xml:"license,omitempty"`
	// 司机姓名
	DriverName string `json:"driver_name,omitempty" xml:"driver_name,omitempty"`
	// 司机联系方式 真实号
	DriverPhone string `json:"driver_phone,omitempty" xml:"driver_phone,omitempty"`
	// 司机推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 行李信息
	Luggage string `json:"luggage,omitempty" xml:"luggage,omitempty"`
}
