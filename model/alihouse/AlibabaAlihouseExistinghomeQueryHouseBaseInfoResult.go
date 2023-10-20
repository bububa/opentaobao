package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeQueryHouseBaseInfoResult 结构体
type AlibabaAlihouseExistinghomeQueryHouseBaseInfoResult struct {
	// 200
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回实体类
	Data *SyncHouseBaseInfoDto `json:"data,omitempty" xml:"data,omitempty"`
	// true
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeQueryHouseBaseInfoResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeQueryHouseBaseInfoResult)
	},
}

// GetAlibabaAlihouseExistinghomeQueryHouseBaseInfoResult() 从对象池中获取AlibabaAlihouseExistinghomeQueryHouseBaseInfoResult
func GetAlibabaAlihouseExistinghomeQueryHouseBaseInfoResult() *AlibabaAlihouseExistinghomeQueryHouseBaseInfoResult {
	return poolAlibabaAlihouseExistinghomeQueryHouseBaseInfoResult.Get().(*AlibabaAlihouseExistinghomeQueryHouseBaseInfoResult)
}

// ReleaseAlibabaAlihouseExistinghomeQueryHouseBaseInfoResult 释放AlibabaAlihouseExistinghomeQueryHouseBaseInfoResult
func ReleaseAlibabaAlihouseExistinghomeQueryHouseBaseInfoResult(v *AlibabaAlihouseExistinghomeQueryHouseBaseInfoResult) {
	v.Code = ""
	v.Message = ""
	v.Data = nil
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeQueryHouseBaseInfoResult.Put(v)
}
