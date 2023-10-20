package baichuan

import (
	"sync"
)

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

var poolAsoDeviceCheckResult = sync.Pool{
	New: func() any {
		return new(AsoDeviceCheckResult)
	},
}

// GetAsoDeviceCheckResult() 从对象池中获取AsoDeviceCheckResult
func GetAsoDeviceCheckResult() *AsoDeviceCheckResult {
	return poolAsoDeviceCheckResult.Get().(*AsoDeviceCheckResult)
}

// ReleaseAsoDeviceCheckResult 释放AsoDeviceCheckResult
func ReleaseAsoDeviceCheckResult(v *AsoDeviceCheckResult) {
	v.Imei = ""
	v.Imsi = ""
	v.Idfa = ""
	v.IsNewDevice = false
	v.IsMyChannal = false
	poolAsoDeviceCheckResult.Put(v)
}
