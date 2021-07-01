package nrt

// AlibabaRetailDeviceVendingRegisterResultDo 结构体
type AlibabaRetailDeviceVendingRegisterResultDo struct {
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
	// 数据
	Data *DeviceDto `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 测试
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
}
