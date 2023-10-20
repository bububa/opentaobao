package ioti

import (
	"sync"
)

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

var poolEslTopEngineAssetsDo = sync.Pool{
	New: func() any {
		return new(EslTopEngineAssetsDo)
	},
}

// GetEslTopEngineAssetsDo() 从对象池中获取EslTopEngineAssetsDo
func GetEslTopEngineAssetsDo() *EslTopEngineAssetsDo {
	return poolEslTopEngineAssetsDo.Get().(*EslTopEngineAssetsDo)
}

// ReleaseEslTopEngineAssetsDo 释放EslTopEngineAssetsDo
func ReleaseEslTopEngineAssetsDo(v *EslTopEngineAssetsDo) {
	v.EslMac = ""
	v.EslModelName = ""
	v.ApMac = ""
	v.Lastseen = 0
	v.FirmwareVersion = 0
	v.BatteryLevel = 0
	v.BatteryVoltage = 0
	v.ApRssi = 0
	poolEslTopEngineAssetsDo.Put(v)
}
