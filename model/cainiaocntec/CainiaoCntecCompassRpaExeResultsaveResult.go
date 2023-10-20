package cainiaocntec

import (
	"sync"
)

// CainiaoCntecCompassRpaExeResultsaveResult 结构体
type CainiaoCntecCompassRpaExeResultsaveResult struct {
	// 请求trace_id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误code
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 请求结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolCainiaoCntecCompassRpaExeResultsaveResult = sync.Pool{
	New: func() any {
		return new(CainiaoCntecCompassRpaExeResultsaveResult)
	},
}

// GetCainiaoCntecCompassRpaExeResultsaveResult() 从对象池中获取CainiaoCntecCompassRpaExeResultsaveResult
func GetCainiaoCntecCompassRpaExeResultsaveResult() *CainiaoCntecCompassRpaExeResultsaveResult {
	return poolCainiaoCntecCompassRpaExeResultsaveResult.Get().(*CainiaoCntecCompassRpaExeResultsaveResult)
}

// ReleaseCainiaoCntecCompassRpaExeResultsaveResult 释放CainiaoCntecCompassRpaExeResultsaveResult
func ReleaseCainiaoCntecCompassRpaExeResultsaveResult(v *CainiaoCntecCompassRpaExeResultsaveResult) {
	v.TraceId = ""
	v.ErrCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolCainiaoCntecCompassRpaExeResultsaveResult.Put(v)
}
