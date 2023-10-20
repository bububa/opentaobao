package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeEntrustsellingQueryResult 结构体
type AlibabaAlihouseExistinghomeEntrustsellingQueryResult struct {
	// 系统自动生成
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data *CustomerEntrustSellingDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeEntrustsellingQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeEntrustsellingQueryResult)
	},
}

// GetAlibabaAlihouseExistinghomeEntrustsellingQueryResult() 从对象池中获取AlibabaAlihouseExistinghomeEntrustsellingQueryResult
func GetAlibabaAlihouseExistinghomeEntrustsellingQueryResult() *AlibabaAlihouseExistinghomeEntrustsellingQueryResult {
	return poolAlibabaAlihouseExistinghomeEntrustsellingQueryResult.Get().(*AlibabaAlihouseExistinghomeEntrustsellingQueryResult)
}

// ReleaseAlibabaAlihouseExistinghomeEntrustsellingQueryResult 释放AlibabaAlihouseExistinghomeEntrustsellingQueryResult
func ReleaseAlibabaAlihouseExistinghomeEntrustsellingQueryResult(v *AlibabaAlihouseExistinghomeEntrustsellingQueryResult) {
	v.Msg = ""
	v.Code = ""
	v.Data = nil
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeEntrustsellingQueryResult.Put(v)
}
