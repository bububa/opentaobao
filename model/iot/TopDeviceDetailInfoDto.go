package iot

// TopDeviceDetailInfoDto 
type TopDeviceDetailInfoDto struct {

    // 设备id
    
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    

    // 设备联网的ip地址
    
    DeviceIp   string `json:"device_ip,omitempty" xml:"device_ip,omitempty"`
    

    // Mac地址
    
    DeviceMac   string `json:"device_mac,omitempty" xml:"device_mac,omitempty"`
    

    // 外部设备id
    
    ExtDeviceId   string `json:"ext_device_id,omitempty" xml:"ext_device_id,omitempty"`
    

    // 外部设备类型，chinaMobile:中国移动，chinaTelecom:上海电信，chinaUnicom:中国联通
    
    ExtDeviceType   string `json:"ext_device_type,omitempty" xml:"ext_device_type,omitempty"`
    

    // 设备当前的固件版本
    
    FirmwareVersion   string `json:"firmware_version,omitempty" xml:"firmware_version,omitempty"`
    

    // 设备序列号
    
    Sn   string `json:"sn,omitempty" xml:"sn,omitempty"`
    

}
