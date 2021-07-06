package baichuan

// AsoDeviceCheckResult 结构体
type AsoDeviceCheckResult struct {
	// imei
	Imei string `json:"imei,omitempty" xml:"imei,omitempty"`
	// imsi
	Imsi string `json:"imsi,omitempty" xml:"imsi,omitempty"`
	// idfa
	Idfa string `json:"idfa,omitempty" xml:"idfa,omitempty"`
	// isNewDevice
	IsNewDevice bool `json:"is_new_device,omitempty" xml:"is_new_device,omitempty"`
	// isMyChannal
	IsMyChannal bool `json:"is_my_channal,omitempty" xml:"is_my_channal,omitempty"`
}
