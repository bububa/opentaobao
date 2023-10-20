package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeHouseDownselfResult 结构体
type AlibabaAlihouseExistinghomeHouseDownselfResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 操作结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseExistinghomeHouseDownselfResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseDownselfResult)
	},
}

// GetAlibabaAlihouseExistinghomeHouseDownselfResult() 从对象池中获取AlibabaAlihouseExistinghomeHouseDownselfResult
func GetAlibabaAlihouseExistinghomeHouseDownselfResult() *AlibabaAlihouseExistinghomeHouseDownselfResult {
	return poolAlibabaAlihouseExistinghomeHouseDownselfResult.Get().(*AlibabaAlihouseExistinghomeHouseDownselfResult)
}

// ReleaseAlibabaAlihouseExistinghomeHouseDownselfResult 释放AlibabaAlihouseExistinghomeHouseDownselfResult
func ReleaseAlibabaAlihouseExistinghomeHouseDownselfResult(v *AlibabaAlihouseExistinghomeHouseDownselfResult) {
	v.Code = ""
	v.Message = ""
	v.IsSuccess = false
	v.Data = false
	poolAlibabaAlihouseExistinghomeHouseDownselfResult.Put(v)
}
