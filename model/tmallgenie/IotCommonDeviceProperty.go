package tmallgenie

// IotCommonDeviceProperty 
type IotCommonDeviceProperty struct {
    // 异常检测项名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 异常检测项值
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
}
