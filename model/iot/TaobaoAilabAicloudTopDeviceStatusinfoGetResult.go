package iot

// TaobaoAilabAicloudTopDeviceStatusinfoGetResult 结构体
type TaobaoAilabAicloudTopDeviceStatusinfoGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 设备状态信息
	Result *TopDeviceStatusInfoDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
