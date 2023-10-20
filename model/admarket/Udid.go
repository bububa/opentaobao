package admarket

import (
	"sync"
)

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

var poolUdid = sync.Pool{
	New: func() any {
		return new(Udid)
	},
}

// GetUdid() 从对象池中获取Udid
func GetUdid() *Udid {
	return poolUdid.Get().(*Udid)
}

// ReleaseUdid 释放Udid
func ReleaseUdid(v *Udid) {
	v.Imei = ""
	v.Mac = ""
	v.AndroidId = ""
	v.UmidToken = ""
	v.Uuid = ""
	v.SerialNum = ""
	v.SimSn = ""
	v.Imsi = ""
	v.Utdid = ""
	v.DeviceId = ""
	poolUdid.Put(v)
}
