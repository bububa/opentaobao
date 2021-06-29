package alilabs

// TopDeviceReqDTO 
type TopDeviceReqDTO struct {
    // 设备签名
    DeviceSignature   string `json:"device_signature,omitempty" xml:"device_signature,omitempty"`
    // 三方设备id
    DeviceId   string `json:"device_id,omitempty" xml:"device_id,omitempty"`
}
