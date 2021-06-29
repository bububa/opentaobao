package iot

// TopDeviceStatusInfoDto 
type TopDeviceStatusInfoDto struct {
    // 设备id
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    // 设备联网的ip地址
    DeviceIp   string `json:"device_ip,omitempty" xml:"device_ip,omitempty"`
    // 设备在线状态，0:离线 1:在线
    OnlineStatus   int64 `json:"online_status,omitempty" xml:"online_status,omitempty"`
}
