package baichuan

import (
	"sync"
)

// AsoDeviceInfoDo 结构体
type AsoDeviceInfoDo struct {
	// imei
	Imei string `json:"imei,omitempty" xml:"imei,omitempty"`
	// imsi
	Imsi string `json:"imsi,omitempty" xml:"imsi,omitempty"`
	// idfa
	Idfa string `json:"idfa,omitempty" xml:"idfa,omitempty"`
}

var poolAsoDeviceInfoDo = sync.Pool{
	New: func() any {
		return new(AsoDeviceInfoDo)
	},
}

// GetAsoDeviceInfoDo() 从对象池中获取AsoDeviceInfoDo
func GetAsoDeviceInfoDo() *AsoDeviceInfoDo {
	return poolAsoDeviceInfoDo.Get().(*AsoDeviceInfoDo)
}

// ReleaseAsoDeviceInfoDo 释放AsoDeviceInfoDo
func ReleaseAsoDeviceInfoDo(v *AsoDeviceInfoDo) {
	v.Imei = ""
	v.Imsi = ""
	v.Idfa = ""
	poolAsoDeviceInfoDo.Put(v)
}
