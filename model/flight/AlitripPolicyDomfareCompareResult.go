package flight

import (
	"sync"
)

// AlitripPolicyDomfareCompareResult 结构体
type AlitripPolicyDomfareCompareResult struct {
	// 调用错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码详情
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 比价结果列表
	Data *CompareDomFareReponseDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripPolicyDomfareCompareResult = sync.Pool{
	New: func() any {
		return new(AlitripPolicyDomfareCompareResult)
	},
}

// GetAlitripPolicyDomfareCompareResult() 从对象池中获取AlitripPolicyDomfareCompareResult
func GetAlitripPolicyDomfareCompareResult() *AlitripPolicyDomfareCompareResult {
	return poolAlitripPolicyDomfareCompareResult.Get().(*AlitripPolicyDomfareCompareResult)
}

// ReleaseAlitripPolicyDomfareCompareResult 释放AlitripPolicyDomfareCompareResult
func ReleaseAlitripPolicyDomfareCompareResult(v *AlitripPolicyDomfareCompareResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolAlitripPolicyDomfareCompareResult.Put(v)
}
