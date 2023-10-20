package tmallsc

import (
	"sync"
)

// AlibabaServicecenterWorkcardRelatedskuQueryResult 结构体
type AlibabaServicecenterWorkcardRelatedskuQueryResult struct {
	// 服务项信息
	Values []Value `json:"values,omitempty" xml:"values>value,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaServicecenterWorkcardRelatedskuQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterWorkcardRelatedskuQueryResult)
	},
}

// GetAlibabaServicecenterWorkcardRelatedskuQueryResult() 从对象池中获取AlibabaServicecenterWorkcardRelatedskuQueryResult
func GetAlibabaServicecenterWorkcardRelatedskuQueryResult() *AlibabaServicecenterWorkcardRelatedskuQueryResult {
	return poolAlibabaServicecenterWorkcardRelatedskuQueryResult.Get().(*AlibabaServicecenterWorkcardRelatedskuQueryResult)
}

// ReleaseAlibabaServicecenterWorkcardRelatedskuQueryResult 释放AlibabaServicecenterWorkcardRelatedskuQueryResult
func ReleaseAlibabaServicecenterWorkcardRelatedskuQueryResult(v *AlibabaServicecenterWorkcardRelatedskuQueryResult) {
	v.Values = v.Values[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolAlibabaServicecenterWorkcardRelatedskuQueryResult.Put(v)
}
