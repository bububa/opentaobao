package iot

import (
	"sync"
)

// TaobaoAilabAicloudTopDeviceDeviceidConvertResult 结构体
type TaobaoAilabAicloudTopDeviceDeviceidConvertResult struct {
	// 描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回成功
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 结果详情
	Result *TopDeviceBaseInfoDto `json:"result,omitempty" xml:"result,omitempty"`
}

var poolTaobaoAilabAicloudTopDeviceDeviceidConvertResult = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceDeviceidConvertResult)
	},
}

// GetTaobaoAilabAicloudTopDeviceDeviceidConvertResult() 从对象池中获取TaobaoAilabAicloudTopDeviceDeviceidConvertResult
func GetTaobaoAilabAicloudTopDeviceDeviceidConvertResult() *TaobaoAilabAicloudTopDeviceDeviceidConvertResult {
	return poolTaobaoAilabAicloudTopDeviceDeviceidConvertResult.Get().(*TaobaoAilabAicloudTopDeviceDeviceidConvertResult)
}

// ReleaseTaobaoAilabAicloudTopDeviceDeviceidConvertResult 释放TaobaoAilabAicloudTopDeviceDeviceidConvertResult
func ReleaseTaobaoAilabAicloudTopDeviceDeviceidConvertResult(v *TaobaoAilabAicloudTopDeviceDeviceidConvertResult) {
	v.Message = ""
	v.Code = 0
	v.Result = nil
	poolTaobaoAilabAicloudTopDeviceDeviceidConvertResult.Put(v)
}
