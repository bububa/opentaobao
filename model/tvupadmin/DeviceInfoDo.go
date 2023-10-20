package tvupadmin

import (
	"sync"
)

// DeviceInfoDo 结构体
type DeviceInfoDo struct {
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// tvid
	Tvid string `json:"tvid,omitempty" xml:"tvid,omitempty"`
	// brandName
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// deviceModel
	DeviceModel string `json:"device_model,omitempty" xml:"device_model,omitempty"`
	// terminalType
	TerminalType string `json:"terminal_type,omitempty" xml:"terminal_type,omitempty"`
	// ethMac
	EthMac string `json:"eth_mac,omitempty" xml:"eth_mac,omitempty"`
	// wlanMac
	WlanMac string `json:"wlan_mac,omitempty" xml:"wlan_mac,omitempty"`
	// activeTime
	ActiveTime string `json:"active_time,omitempty" xml:"active_time,omitempty"`
	// userActiveTime
	UserActiveTime string `json:"user_active_time,omitempty" xml:"user_active_time,omitempty"`
	// lastLoginTime
	LastLoginTime string `json:"last_login_time,omitempty" xml:"last_login_time,omitempty"`
	// createdAt
	CreatedAt string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// updatedAt
	UpdatedAt string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// systemVersion
	SystemVersion string `json:"system_version,omitempty" xml:"system_version,omitempty"`
	// brandId
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
}

var poolDeviceInfoDo = sync.Pool{
	New: func() any {
		return new(DeviceInfoDo)
	},
}

// GetDeviceInfoDo() 从对象池中获取DeviceInfoDo
func GetDeviceInfoDo() *DeviceInfoDo {
	return poolDeviceInfoDo.Get().(*DeviceInfoDo)
}

// ReleaseDeviceInfoDo 释放DeviceInfoDo
func ReleaseDeviceInfoDo(v *DeviceInfoDo) {
	v.Uuid = ""
	v.Tvid = ""
	v.BrandName = ""
	v.DeviceModel = ""
	v.TerminalType = ""
	v.EthMac = ""
	v.WlanMac = ""
	v.ActiveTime = ""
	v.UserActiveTime = ""
	v.LastLoginTime = ""
	v.CreatedAt = ""
	v.UpdatedAt = ""
	v.SystemVersion = ""
	v.BrandId = 0
	poolDeviceInfoDo.Put(v)
}
