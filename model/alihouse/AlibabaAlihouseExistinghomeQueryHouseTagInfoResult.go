package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeQueryHouseTagInfoResult 结构体
type AlibabaAlihouseExistinghomeQueryHouseTagInfoResult struct {
	// 1
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 1
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回值
	Data *SyncHouseTagFeatureDto `json:"data,omitempty" xml:"data,omitempty"`
	// 1
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeQueryHouseTagInfoResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeQueryHouseTagInfoResult)
	},
}

// GetAlibabaAlihouseExistinghomeQueryHouseTagInfoResult() 从对象池中获取AlibabaAlihouseExistinghomeQueryHouseTagInfoResult
func GetAlibabaAlihouseExistinghomeQueryHouseTagInfoResult() *AlibabaAlihouseExistinghomeQueryHouseTagInfoResult {
	return poolAlibabaAlihouseExistinghomeQueryHouseTagInfoResult.Get().(*AlibabaAlihouseExistinghomeQueryHouseTagInfoResult)
}

// ReleaseAlibabaAlihouseExistinghomeQueryHouseTagInfoResult 释放AlibabaAlihouseExistinghomeQueryHouseTagInfoResult
func ReleaseAlibabaAlihouseExistinghomeQueryHouseTagInfoResult(v *AlibabaAlihouseExistinghomeQueryHouseTagInfoResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeQueryHouseTagInfoResult.Put(v)
}
