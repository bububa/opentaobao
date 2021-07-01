package iot

// TaobaoAilabAicloudTopDeviceExtinfoGetResult 结构体
type TaobaoAilabAicloudTopDeviceExtinfoGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 三方设备信息
	Result *TopDeviceExtInfoDto `json:"result,omitempty" xml:"result,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
