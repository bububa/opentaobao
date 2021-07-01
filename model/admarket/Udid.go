package admarket

// Udid 结构体
type Udid struct {
	// imei
	Imei string `json:"imei,omitempty" xml:"imei,omitempty"`
	// mac
	Mac string `json:"mac,omitempty" xml:"mac,omitempty"`
	// andorid设备id
	AndroidId string `json:"android_id,omitempty" xml:"android_id,omitempty"`
	// umidToken
	UmidToken string `json:"umid_token,omitempty" xml:"umid_token,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 设备序列号
	SerialNum string `json:"serial_num,omitempty" xml:"serial_num,omitempty"`
	// sim序列号
	SimSn string `json:"sim_sn,omitempty" xml:"sim_sn,omitempty"`
	// 设备序列号
	Imsi string `json:"imsi,omitempty" xml:"imsi,omitempty"`
	// utdid
	Utdid string `json:"utdid,omitempty" xml:"utdid,omitempty"`
	// 设备id
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
}
