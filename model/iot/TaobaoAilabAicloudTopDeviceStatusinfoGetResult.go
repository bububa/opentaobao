package iot

import (
	"sync"
)

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

var poolTaobaoAilabAicloudTopDeviceStatusinfoGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopDeviceStatusinfoGetResult)
	},
}

// GetTaobaoAilabAicloudTopDeviceStatusinfoGetResult() 从对象池中获取TaobaoAilabAicloudTopDeviceStatusinfoGetResult
func GetTaobaoAilabAicloudTopDeviceStatusinfoGetResult() *TaobaoAilabAicloudTopDeviceStatusinfoGetResult {
	return poolTaobaoAilabAicloudTopDeviceStatusinfoGetResult.Get().(*TaobaoAilabAicloudTopDeviceStatusinfoGetResult)
}

// ReleaseTaobaoAilabAicloudTopDeviceStatusinfoGetResult 释放TaobaoAilabAicloudTopDeviceStatusinfoGetResult
func ReleaseTaobaoAilabAicloudTopDeviceStatusinfoGetResult(v *TaobaoAilabAicloudTopDeviceStatusinfoGetResult) {
	v.Message = ""
	v.Code = 0
	v.Result = nil
	v.Success = false
	poolTaobaoAilabAicloudTopDeviceStatusinfoGetResult.Put(v)
}
