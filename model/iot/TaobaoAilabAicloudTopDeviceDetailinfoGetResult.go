package iot

// TaobaoAilabAicloudTopDeviceDetailinfoGetResult 结构体
type TaobaoAilabAicloudTopDeviceDetailinfoGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 设备详细信息
	Result *TopDeviceDetailInfoDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
