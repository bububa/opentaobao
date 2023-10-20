package iot

import (
	"sync"
)

// TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult 结构体
type TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// exception
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// statusCode
	StatusCode int64 `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// result
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult)
	},
}

// GetTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult() 从对象池中获取TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult
func GetTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult() *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult {
	return poolTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult.Get().(*TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult)
}

// ReleaseTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult 释放TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult
func ReleaseTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult(v *TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult) {
	v.Message = ""
	v.Exception = ""
	v.StatusCode = 0
	v.Result = false
	poolTaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceResult.Put(v)
}
