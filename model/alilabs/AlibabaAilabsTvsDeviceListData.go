package alilabs

// AlibabaAilabsTvsDeviceListData 结构体
type AlibabaAilabsTvsDeviceListData struct {
	// 设备名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 设备激活时间
	ActivateTime string `json:"activate_time,omitempty" xml:"activate_time,omitempty"`
	// TVS服务给予设备的唯一标识
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 厂商给定的设备硬件标识
	Dsn string `json:"dsn,omitempty" xml:"dsn,omitempty"`
}
