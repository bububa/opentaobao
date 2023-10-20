package alihouse

import (
	"sync"
)

// AlibabaAlihouseNewhomeLayoutEcodeUpdateResult 结构体
type AlibabaAlihouseNewhomeLayoutEcodeUpdateResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回户型ID
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseNewhomeLayoutEcodeUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeLayoutEcodeUpdateResult)
	},
}

// GetAlibabaAlihouseNewhomeLayoutEcodeUpdateResult() 从对象池中获取AlibabaAlihouseNewhomeLayoutEcodeUpdateResult
func GetAlibabaAlihouseNewhomeLayoutEcodeUpdateResult() *AlibabaAlihouseNewhomeLayoutEcodeUpdateResult {
	return poolAlibabaAlihouseNewhomeLayoutEcodeUpdateResult.Get().(*AlibabaAlihouseNewhomeLayoutEcodeUpdateResult)
}

// ReleaseAlibabaAlihouseNewhomeLayoutEcodeUpdateResult 释放AlibabaAlihouseNewhomeLayoutEcodeUpdateResult
func ReleaseAlibabaAlihouseNewhomeLayoutEcodeUpdateResult(v *AlibabaAlihouseNewhomeLayoutEcodeUpdateResult) {
	v.Message = ""
	v.Code = ""
	v.Data = ""
	v.Success = false
	poolAlibabaAlihouseNewhomeLayoutEcodeUpdateResult.Put(v)
}
