package flight

import (
	"sync"
)

// AlitripPolicyDomfareFlowdataResult 结构体
type AlitripPolicyDomfareFlowdataResult struct {
	// 返回错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回的错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回的信息
	Data *CompareFlowDataReponseDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPolicyDomfareFlowdataResult = sync.Pool{
	New: func() any {
		return new(AlitripPolicyDomfareFlowdataResult)
	},
}

// GetAlitripPolicyDomfareFlowdataResult() 从对象池中获取AlitripPolicyDomfareFlowdataResult
func GetAlitripPolicyDomfareFlowdataResult() *AlitripPolicyDomfareFlowdataResult {
	return poolAlitripPolicyDomfareFlowdataResult.Get().(*AlitripPolicyDomfareFlowdataResult)
}

// ReleaseAlitripPolicyDomfareFlowdataResult 释放AlitripPolicyDomfareFlowdataResult
func ReleaseAlitripPolicyDomfareFlowdataResult(v *AlitripPolicyDomfareFlowdataResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolAlitripPolicyDomfareFlowdataResult.Put(v)
}
