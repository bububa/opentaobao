package campus

// DeviceReportDataDto 结构体
type DeviceReportDataDto struct {
	// 温度
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 描述信息
	Describe string `json:"describe,omitempty" xml:"describe,omitempty"`
	// 值
	OriginValue string `json:"origin_value,omitempty" xml:"origin_value,omitempty"`
	// 是否告警
	IsAlarm bool `json:"is_alarm,omitempty" xml:"is_alarm,omitempty"`
}
