package alsc

import (
	"sync"
)

// DeviceInfo 结构体
type DeviceInfo struct {
	// 设备ID
	DeviceId string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	// ip地址
	Ip string `json:"ip,omitempty" xml:"ip,omitempty"`
	// 设备MAC地址
	Mac string `json:"mac,omitempty" xml:"mac,omitempty"`
	// 客户端平台
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
	// 客户端平台版本
	PlatformVersion string `json:"platform_version,omitempty" xml:"platform_version,omitempty"`
	// 客户端应用版本
	ProdVersion string `json:"prod_version,omitempty" xml:"prod_version,omitempty"`
	// 客户端应用(淘宝、口碑独客、支付宝tab…)
	Product string `json:"product,omitempty" xml:"product,omitempty"`
	// 新零售店铺跳转链接:透传给投承做补全新零售店铺链接时使用
	SchemaPlatform string `json:"schema_platform,omitempty" xml:"schema_platform,omitempty"`
	// 格式: xxx@taobao_ios_5.0.1
	Ttid string `json:"ttid,omitempty" xml:"ttid,omitempty"`
	// 特征字符串
	UserAgent string `json:"user_agent,omitempty" xml:"user_agent,omitempty"`
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
	v.DeviceId = ""
	v.Ip = ""
	v.Mac = ""
	v.Platform = ""
	v.PlatformVersion = ""
	v.ProdVersion = ""
	v.Product = ""
	v.SchemaPlatform = ""
	v.Ttid = ""
	v.UserAgent = ""
	poolDeviceInfo.Put(v)
}
