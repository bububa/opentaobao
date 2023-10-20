package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeHouseChangeStandardResult 结构体
type AlibabaAlihouseExistinghomeHouseChangeStandardResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 操作结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseExistinghomeHouseChangeStandardResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseChangeStandardResult)
	},
}

// GetAlibabaAlihouseExistinghomeHouseChangeStandardResult() 从对象池中获取AlibabaAlihouseExistinghomeHouseChangeStandardResult
func GetAlibabaAlihouseExistinghomeHouseChangeStandardResult() *AlibabaAlihouseExistinghomeHouseChangeStandardResult {
	return poolAlibabaAlihouseExistinghomeHouseChangeStandardResult.Get().(*AlibabaAlihouseExistinghomeHouseChangeStandardResult)
}

// ReleaseAlibabaAlihouseExistinghomeHouseChangeStandardResult 释放AlibabaAlihouseExistinghomeHouseChangeStandardResult
func ReleaseAlibabaAlihouseExistinghomeHouseChangeStandardResult(v *AlibabaAlihouseExistinghomeHouseChangeStandardResult) {
	v.Code = ""
	v.Message = ""
	v.IsSuccess = false
	v.Data = false
	poolAlibabaAlihouseExistinghomeHouseChangeStandardResult.Put(v)
}
