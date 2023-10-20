package ihome

import (
	"sync"
)

// AlibabaIhomeCtomCaseMainpicUpdateApiResult 结构体
type AlibabaIhomeCtomCaseMainpicUpdateApiResult struct {
	// 具体错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// case的url地址
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// true
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}

var poolAlibabaIhomeCtomCaseMainpicUpdateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaIhomeCtomCaseMainpicUpdateApiResult)
	},
}

// GetAlibabaIhomeCtomCaseMainpicUpdateApiResult() 从对象池中获取AlibabaIhomeCtomCaseMainpicUpdateApiResult
func GetAlibabaIhomeCtomCaseMainpicUpdateApiResult() *AlibabaIhomeCtomCaseMainpicUpdateApiResult {
	return poolAlibabaIhomeCtomCaseMainpicUpdateApiResult.Get().(*AlibabaIhomeCtomCaseMainpicUpdateApiResult)
}

// ReleaseAlibabaIhomeCtomCaseMainpicUpdateApiResult 释放AlibabaIhomeCtomCaseMainpicUpdateApiResult
func ReleaseAlibabaIhomeCtomCaseMainpicUpdateApiResult(v *AlibabaIhomeCtomCaseMainpicUpdateApiResult) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Data = ""
	v.Status = false
	poolAlibabaIhomeCtomCaseMainpicUpdateApiResult.Put(v)
}
