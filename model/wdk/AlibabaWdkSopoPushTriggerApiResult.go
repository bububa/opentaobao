package wdk

import (
	"sync"
)

// AlibabaWdkSopoPushTriggerApiResult 结构体
type AlibabaWdkSopoPushTriggerApiResult struct {
	// 错误码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 错误信息
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 结果
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}

var poolAlibabaWdkSopoPushTriggerApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSopoPushTriggerApiResult)
	},
}

// GetAlibabaWdkSopoPushTriggerApiResult() 从对象池中获取AlibabaWdkSopoPushTriggerApiResult
func GetAlibabaWdkSopoPushTriggerApiResult() *AlibabaWdkSopoPushTriggerApiResult {
	return poolAlibabaWdkSopoPushTriggerApiResult.Get().(*AlibabaWdkSopoPushTriggerApiResult)
}

// ReleaseAlibabaWdkSopoPushTriggerApiResult 释放AlibabaWdkSopoPushTriggerApiResult
func ReleaseAlibabaWdkSopoPushTriggerApiResult(v *AlibabaWdkSopoPushTriggerApiResult) {
	v.ReturnCode = ""
	v.ReturnMsg = ""
	v.ReturnSuccess = false
	poolAlibabaWdkSopoPushTriggerApiResult.Put(v)
}
