package nrt

// DeviceDto 
type DeviceDto struct {
    // 设备编码
    DeviceCode   string `json:"device_code,omitempty" xml:"device_code,omitempty"`
    // 设备外部编码
    DeviceUuid   string `json:"device_uuid,omitempty" xml:"device_uuid,omitempty"`
    // 设备物理编码
    DeviceSn   string `json:"device_sn,omitempty" xml:"device_sn,omitempty"`
    // 设备状态0离线，1 在线 2 待激活
    DeviceStatus   string `json:"device_status,omitempty" xml:"device_status,omitempty"`
    // 设备名称
    DeviceName   string `json:"device_name,omitempty" xml:"device_name,omitempty"`
    // 设备地址
    DeviceAddress   string `json:"device_address,omitempty" xml:"device_address,omitempty"`
    // 设备类型
    DeviceTypeName   string `json:"device_type_name,omitempty" xml:"device_type_name,omitempty"`
    // 经度
    Longtitude   string `json:"longtitude,omitempty" xml:"longtitude,omitempty"`
    // 纬度
    Latitude   string `json:"latitude,omitempty" xml:"latitude,omitempty"`
    // 激活码
    Activation   string `json:"activation,omitempty" xml:"activation,omitempty"`
}
