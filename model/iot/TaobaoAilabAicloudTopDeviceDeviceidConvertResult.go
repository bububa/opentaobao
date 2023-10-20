package iot

// TaobaoailabaicloudtopdevicedeviceidconvertResult 结构体
type TaobaoailabaicloudtopdevicedeviceidconvertResult struct {
	// 描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回成功
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 结果详情
	Result *TopDeviceBaseInfoDto `json:"result,omitempty" xml:"result,omitempty"`
}
