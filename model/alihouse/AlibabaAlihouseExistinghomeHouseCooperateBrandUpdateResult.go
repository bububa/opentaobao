package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult 结构体
type AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult struct {
	// 返回编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回具体值
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 返回成功失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult)
	},
}

// GetAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult() 从对象池中获取AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult
func GetAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult() *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult {
	return poolAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult.Get().(*AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult)
}

// ReleaseAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult 释放AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult
func ReleaseAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult(v *AlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeHouseCooperateBrandUpdateResult.Put(v)
}
