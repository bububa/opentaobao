package util

// IotRegisterResult 结构体
type IotRegisterResult struct {
	// 产品Key
	ProductKey string `json:"product_key,omitempty" xml:"product_key,omitempty"`
	// 设备名称
	DeviceName string `json:"device_name,omitempty" xml:"device_name,omitempty"`
	// 设备Secret
	DeviceSecret string `json:"device_secret,omitempty" xml:"device_secret,omitempty"`
}
