package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeBrokerQueryResult 结构体
type AlibabaAlihouseExistinghomeBrokerQueryResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data *ProjectAdviserDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeBrokerQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeBrokerQueryResult)
	},
}

// GetAlibabaAlihouseExistinghomeBrokerQueryResult() 从对象池中获取AlibabaAlihouseExistinghomeBrokerQueryResult
func GetAlibabaAlihouseExistinghomeBrokerQueryResult() *AlibabaAlihouseExistinghomeBrokerQueryResult {
	return poolAlibabaAlihouseExistinghomeBrokerQueryResult.Get().(*AlibabaAlihouseExistinghomeBrokerQueryResult)
}

// ReleaseAlibabaAlihouseExistinghomeBrokerQueryResult 释放AlibabaAlihouseExistinghomeBrokerQueryResult
func ReleaseAlibabaAlihouseExistinghomeBrokerQueryResult(v *AlibabaAlihouseExistinghomeBrokerQueryResult) {
	v.Message = ""
	v.Code = ""
	v.Data = nil
	v.Success = false
	poolAlibabaAlihouseExistinghomeBrokerQueryResult.Put(v)
}
