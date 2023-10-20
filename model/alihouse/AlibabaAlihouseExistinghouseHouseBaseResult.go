package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghouseHouseBaseResult 结构体
type AlibabaAlihouseExistinghouseHouseBaseResult struct {
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果集
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihouseExistinghouseHouseBaseResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghouseHouseBaseResult)
	},
}

// GetAlibabaAlihouseExistinghouseHouseBaseResult() 从对象池中获取AlibabaAlihouseExistinghouseHouseBaseResult
func GetAlibabaAlihouseExistinghouseHouseBaseResult() *AlibabaAlihouseExistinghouseHouseBaseResult {
	return poolAlibabaAlihouseExistinghouseHouseBaseResult.Get().(*AlibabaAlihouseExistinghouseHouseBaseResult)
}

// ReleaseAlibabaAlihouseExistinghouseHouseBaseResult 释放AlibabaAlihouseExistinghouseHouseBaseResult
func ReleaseAlibabaAlihouseExistinghouseHouseBaseResult(v *AlibabaAlihouseExistinghouseHouseBaseResult) {
	v.Code = ""
	v.Message = ""
	v.Data = 0
	v.Success = false
	poolAlibabaAlihouseExistinghouseHouseBaseResult.Put(v)
}
