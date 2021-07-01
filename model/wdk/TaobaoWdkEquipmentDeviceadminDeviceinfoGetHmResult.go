package wdk

// TaobaoWdkEquipmentDeviceadminDeviceinfoGetHmResult 结构体
type TaobaoWdkEquipmentDeviceadminDeviceinfoGetHmResult struct {
	// 设备列表
	Models []DeviceInfoDto `json:"models,omitempty" xml:"models>device_info_dto,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
