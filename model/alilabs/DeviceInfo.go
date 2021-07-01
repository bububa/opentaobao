package alilabs

// DeviceInfo 结构体
type DeviceInfo struct {
	// 设备唯一id
	DevId string `json:"dev_id,omitempty" xml:"dev_id,omitempty"`
	// 设备状态Map，key和value均为String
	Status *Status `json:"status,omitempty" xml:"status,omitempty"`
}
