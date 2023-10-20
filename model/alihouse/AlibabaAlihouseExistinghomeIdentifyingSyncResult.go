package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeIdentifyingSyncResult 结构体
type AlibabaAlihouseExistinghomeIdentifyingSyncResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 处理结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// message
	Message bool `json:"message,omitempty" xml:"message,omitempty"`
}

var poolAlibabaAlihouseExistinghomeIdentifyingSyncResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeIdentifyingSyncResult)
	},
}

// GetAlibabaAlihouseExistinghomeIdentifyingSyncResult() 从对象池中获取AlibabaAlihouseExistinghomeIdentifyingSyncResult
func GetAlibabaAlihouseExistinghomeIdentifyingSyncResult() *AlibabaAlihouseExistinghomeIdentifyingSyncResult {
	return poolAlibabaAlihouseExistinghomeIdentifyingSyncResult.Get().(*AlibabaAlihouseExistinghomeIdentifyingSyncResult)
}

// ReleaseAlibabaAlihouseExistinghomeIdentifyingSyncResult 释放AlibabaAlihouseExistinghomeIdentifyingSyncResult
func ReleaseAlibabaAlihouseExistinghomeIdentifyingSyncResult(v *AlibabaAlihouseExistinghomeIdentifyingSyncResult) {
	v.Code = ""
	v.Data = false
	v.Success = false
	v.Message = false
	poolAlibabaAlihouseExistinghomeIdentifyingSyncResult.Put(v)
}
