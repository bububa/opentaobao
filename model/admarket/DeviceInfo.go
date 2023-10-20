package admarket

import (
	"sync"
)

// DeviceInfo 结构体
type DeviceInfo struct {
	// 机型
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 操作系统版本
	OsVersion string `json:"os_version,omitempty" xml:"os_version,omitempty"`
	// 设备类型(OFFLINE_MACHINE/PHONE等)
	DeviceType string `json:"device_type,omitempty" xml:"device_type,omitempty"`
	// 厂商名称
	Vendor string `json:"vendor,omitempty" xml:"vendor,omitempty"`
	// 操作系统类型(ANDROID/ALIOS/WINDOWS/IOS)
	OsType string `json:"os_type,omitempty" xml:"os_type,omitempty"`
	// 客户端类型(MULTIPORT_SDK/PHONE_SDK)
	ClientType string `json:"client_type,omitempty" xml:"client_type,omitempty"`
	// 屏幕类型(VENDING_MACHINE)
	ScreenType string `json:"screen_type,omitempty" xml:"screen_type,omitempty"`
	// 屏幕分辨率宽度
	ScreenWidth int64 `json:"screen_width,omitempty" xml:"screen_width,omitempty"`
	// 分辨率高
	ScreenHeight int64 `json:"screen_height,omitempty" xml:"screen_height,omitempty"`
}

var poolDeviceInfo = sync.Pool{
	New: func() any {
		return new(DeviceInfo)
	},
}

// GetDeviceInfo() 从对象池中获取DeviceInfo
func GetDeviceInfo() *DeviceInfo {
	return poolDeviceInfo.Get().(*DeviceInfo)
}

// ReleaseDeviceInfo 释放DeviceInfo
func ReleaseDeviceInfo(v *DeviceInfo) {
	v.Model = ""
	v.OsVersion = ""
	v.DeviceType = ""
	v.Vendor = ""
	v.OsType = ""
	v.ClientType = ""
	v.ScreenType = ""
	v.ScreenWidth = 0
	v.ScreenHeight = 0
	poolDeviceInfo.Put(v)
}
