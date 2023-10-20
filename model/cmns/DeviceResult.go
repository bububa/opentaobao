package cmns

import (
	"sync"
)

// DeviceResult 结构体
type DeviceResult struct {
	// 激活时间
	ActiveTime string `json:"active_time,omitempty" xml:"active_time,omitempty"`
	// 基带
	BaseType string `json:"base_type,omitempty" xml:"base_type,omitempty"`
	// BSP
	BspType string `json:"bsp_type,omitempty" xml:"bsp_type,omitempty"`
	// 有线网卡mac
	EthMac string `json:"eth_mac,omitempty" xml:"eth_mac,omitempty"`
	// imei
	Imei string `json:"imei,omitempty" xml:"imei,omitempty"`
	// imsi1
	Imsi1 string `json:"imsi1,omitempty" xml:"imsi1,omitempty"`
	// imsi2
	Imsi2 string `json:"imsi2,omitempty" xml:"imsi2,omitempty"`
	// isVendorRom
	IsVendorRom string `json:"is_vendor_rom,omitempty" xml:"is_vendor_rom,omitempty"`
	// kp
	Kp string `json:"kp,omitempty" xml:"kp,omitempty"`
	// 网络类型
	Networking string `json:"networking,omitempty" xml:"networking,omitempty"`
	// phoneType
	PhoneType string `json:"phone_type,omitempty" xml:"phone_type,omitempty"`
	// 系统版本
	Rom string `json:"rom,omitempty" xml:"rom,omitempty"`
	// romTuiguang
	RomTuiguang string `json:"rom_tuiguang,omitempty" xml:"rom_tuiguang,omitempty"`
	// romVendor
	RomVendor string `json:"rom_vendor,omitempty" xml:"rom_vendor,omitempty"`
	// 设备机型
	Terminal string `json:"terminal,omitempty" xml:"terminal,omitempty"`
	// 轻量升级版本
	UpdateVersion string `json:"update_version,omitempty" xml:"update_version,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 无线mac
	WlanMac string `json:"wlan_mac,omitempty" xml:"wlan_mac,omitempty"`
	// 是否接收消息
	AcceptMessage int64 `json:"accept_message,omitempty" xml:"accept_message,omitempty"`
	// 设备类型：1手机 2TV
	DeviceType int64 `json:"device_type,omitempty" xml:"device_type,omitempty"`
	// did
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 最近登录时间
	LoginTime int64 `json:"login_time,omitempty" xml:"login_time,omitempty"`
	// 是否在线,1为在线，0为离线
	Online int64 `json:"online,omitempty" xml:"online,omitempty"`
}

var poolDeviceResult = sync.Pool{
	New: func() any {
		return new(DeviceResult)
	},
}

// GetDeviceResult() 从对象池中获取DeviceResult
func GetDeviceResult() *DeviceResult {
	return poolDeviceResult.Get().(*DeviceResult)
}

// ReleaseDeviceResult 释放DeviceResult
func ReleaseDeviceResult(v *DeviceResult) {
	v.ActiveTime = ""
	v.BaseType = ""
	v.BspType = ""
	v.EthMac = ""
	v.Imei = ""
	v.Imsi1 = ""
	v.Imsi2 = ""
	v.IsVendorRom = ""
	v.Kp = ""
	v.Networking = ""
	v.PhoneType = ""
	v.Rom = ""
	v.RomTuiguang = ""
	v.RomVendor = ""
	v.Terminal = ""
	v.UpdateVersion = ""
	v.Uuid = ""
	v.WlanMac = ""
	v.AcceptMessage = 0
	v.DeviceType = 0
	v.Id = 0
	v.LoginTime = 0
	v.Online = 0
	poolDeviceResult.Put(v)
}
