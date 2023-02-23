package ioti

// EslTopEngineAssetsDo 结构体
type EslTopEngineAssetsDo struct {
	// mac
	EslMac string `json:"esl_mac,omitempty" xml:"esl_mac,omitempty"`
	// 价签型号
	EslModelName string `json:"esl_model_name,omitempty" xml:"esl_model_name,omitempty"`
	// 上传来的ap
	ApMac string `json:"ap_mac,omitempty" xml:"ap_mac,omitempty"`
	// 最后上报时间
	Lastseen int64 `json:"lastseen,omitempty" xml:"lastseen,omitempty"`
	// 固件型号
	FirmwareVersion int64 `json:"firmware_version,omitempty" xml:"firmware_version,omitempty"`
	// 电池电量
	BatteryLevel int64 `json:"battery_level,omitempty" xml:"battery_level,omitempty"`
	// 电池电压
	BatteryVoltage int64 `json:"battery_voltage,omitempty" xml:"battery_voltage,omitempty"`
	// ap上次上报的最强rssi值
	ApRssi int64 `json:"ap_rssi,omitempty" xml:"ap_rssi,omitempty"`
}
