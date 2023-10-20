package alihealth2

import (
	"sync"
)

// AlibabaAlihealthTracecodeplatformCodeActiveResult 结构体
type AlibabaAlihealthTracecodeplatformCodeActiveResult struct {
	// 状态代码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 状态信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 调用状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthTracecodeplatformCodeActiveResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodeplatformCodeActiveResult)
	},
}

// GetAlibabaAlihealthTracecodeplatformCodeActiveResult() 从对象池中获取AlibabaAlihealthTracecodeplatformCodeActiveResult
func GetAlibabaAlihealthTracecodeplatformCodeActiveResult() *AlibabaAlihealthTracecodeplatformCodeActiveResult {
	return poolAlibabaAlihealthTracecodeplatformCodeActiveResult.Get().(*AlibabaAlihealthTracecodeplatformCodeActiveResult)
}

// ReleaseAlibabaAlihealthTracecodeplatformCodeActiveResult 释放AlibabaAlihealthTracecodeplatformCodeActiveResult
func ReleaseAlibabaAlihealthTracecodeplatformCodeActiveResult(v *AlibabaAlihealthTracecodeplatformCodeActiveResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolAlibabaAlihealthTracecodeplatformCodeActiveResult.Put(v)
}
