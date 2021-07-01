package baichuan

// AsoDeviceInfoDo 结构体
type AsoDeviceInfoDo struct {
	// imei
	Imei string `json:"imei,omitempty" xml:"imei,omitempty"`
	// imsi
	Imsi string `json:"imsi,omitempty" xml:"imsi,omitempty"`
	// idfa
	Idfa string `json:"idfa,omitempty" xml:"idfa,omitempty"`
}
