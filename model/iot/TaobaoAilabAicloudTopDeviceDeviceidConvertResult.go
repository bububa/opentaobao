package iot

// TaobaoAilabAicloudTopDeviceDeviceidConvertResult 结构体
type TaobaoAilabAicloudTopDeviceDeviceidConvertResult struct {
	// 返回成功
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果详情
	Result *TopDeviceBaseInfoDto `json:"result,omitempty" xml:"result,omitempty"`
}
