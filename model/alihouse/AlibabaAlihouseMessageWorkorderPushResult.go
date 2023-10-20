package alihouse

import (
	"sync"
)

// AlibabaAlihouseMessageWorkorderPushResult 结构体
type AlibabaAlihouseMessageWorkorderPushResult struct {
	// 错误信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 消息记录id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseMessageWorkorderPushResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseMessageWorkorderPushResult)
	},
}

// GetAlibabaAlihouseMessageWorkorderPushResult() 从对象池中获取AlibabaAlihouseMessageWorkorderPushResult
func GetAlibabaAlihouseMessageWorkorderPushResult() *AlibabaAlihouseMessageWorkorderPushResult {
	return poolAlibabaAlihouseMessageWorkorderPushResult.Get().(*AlibabaAlihouseMessageWorkorderPushResult)
}

// ReleaseAlibabaAlihouseMessageWorkorderPushResult 释放AlibabaAlihouseMessageWorkorderPushResult
func ReleaseAlibabaAlihouseMessageWorkorderPushResult(v *AlibabaAlihouseMessageWorkorderPushResult) {
	v.Msg = ""
	v.Code = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseMessageWorkorderPushResult.Put(v)
}
