package iot

import (
	"sync"
)

// TaobaoAilabAicloudTopDeviceExtinfoGetResult 结构体
type TaobaoAilabAicloudTopDeviceExtinfoGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 三方设备信息
	Result *TopDeviceExtInfoDto `json:"result,omitempty" xml:"result,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolTaobaoAilabAicloudTopDeviceExtinfoGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceExtinfoGetResult)
	},
}

// GetTaobaoAilabAicloudTopDeviceExtinfoGetResult() 从对象池中获取TaobaoAilabAicloudTopDeviceExtinfoGetResult
func GetTaobaoAilabAicloudTopDeviceExtinfoGetResult() *TaobaoAilabAicloudTopDeviceExtinfoGetResult {
	return poolTaobaoAilabAicloudTopDeviceExtinfoGetResult.Get().(*TaobaoAilabAicloudTopDeviceExtinfoGetResult)
}

// ReleaseTaobaoAilabAicloudTopDeviceExtinfoGetResult 释放TaobaoAilabAicloudTopDeviceExtinfoGetResult
func ReleaseTaobaoAilabAicloudTopDeviceExtinfoGetResult(v *TaobaoAilabAicloudTopDeviceExtinfoGetResult) {
	v.Message = ""
	v.Result = nil
	v.Code = 0
	poolTaobaoAilabAicloudTopDeviceExtinfoGetResult.Put(v)
}
