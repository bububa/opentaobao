package iot

// TopDeviceExtInfoDto 
type TopDeviceExtInfoDto struct {

    // 设备id
    
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
    

    // 三方设备id
    
    ExtDeviceId   string `json:"ext_device_id,omitempty" xml:"ext_device_id,omitempty"`
    

    // 三方设备类型
    
    ExtDeviceType   string `json:"ext_device_type,omitempty" xml:"ext_device_type,omitempty"`
    

    // 设备mac
    
    DeviceMac   string `json:"device_mac,omitempty" xml:"device_mac,omitempty"`
    

}
