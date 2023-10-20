package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeHouseDeleteHouseResult 结构体
type AlibabaAlihouseExistinghomeHouseDeleteHouseResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 操作结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

var poolAlibabaAlihouseExistinghomeHouseDeleteHouseResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseDeleteHouseResult)
	},
}

// GetAlibabaAlihouseExistinghomeHouseDeleteHouseResult() 从对象池中获取AlibabaAlihouseExistinghomeHouseDeleteHouseResult
func GetAlibabaAlihouseExistinghomeHouseDeleteHouseResult() *AlibabaAlihouseExistinghomeHouseDeleteHouseResult {
	return poolAlibabaAlihouseExistinghomeHouseDeleteHouseResult.Get().(*AlibabaAlihouseExistinghomeHouseDeleteHouseResult)
}

// ReleaseAlibabaAlihouseExistinghomeHouseDeleteHouseResult 释放AlibabaAlihouseExistinghomeHouseDeleteHouseResult
func ReleaseAlibabaAlihouseExistinghomeHouseDeleteHouseResult(v *AlibabaAlihouseExistinghomeHouseDeleteHouseResult) {
	v.Code = ""
	v.Message = ""
	v.IsSuccess = false
	v.Data = false
	poolAlibabaAlihouseExistinghomeHouseDeleteHouseResult.Put(v)
}
