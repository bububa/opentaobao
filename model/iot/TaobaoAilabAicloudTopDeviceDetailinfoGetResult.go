package iot

import (
	"sync"
)

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

var poolTaobaoAilabAicloudTopDeviceDetailinfoGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceDetailinfoGetResult)
	},
}

// GetTaobaoAilabAicloudTopDeviceDetailinfoGetResult() 从对象池中获取TaobaoAilabAicloudTopDeviceDetailinfoGetResult
func GetTaobaoAilabAicloudTopDeviceDetailinfoGetResult() *TaobaoAilabAicloudTopDeviceDetailinfoGetResult {
	return poolTaobaoAilabAicloudTopDeviceDetailinfoGetResult.Get().(*TaobaoAilabAicloudTopDeviceDetailinfoGetResult)
}

// ReleaseTaobaoAilabAicloudTopDeviceDetailinfoGetResult 释放TaobaoAilabAicloudTopDeviceDetailinfoGetResult
func ReleaseTaobaoAilabAicloudTopDeviceDetailinfoGetResult(v *TaobaoAilabAicloudTopDeviceDetailinfoGetResult) {
	v.Message = ""
	v.Code = 0
	v.Result = nil
	v.Success = false
	poolTaobaoAilabAicloudTopDeviceDetailinfoGetResult.Put(v)
}
