package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeHouseStoreentrustQueryResult 结构体
type AlibabaAlihouseExistinghomeHouseStoreentrustQueryResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回素材id
	Data *StoreEntrustDto `json:"data,omitempty" xml:"data,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeHouseStoreentrustQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseStoreentrustQueryResult)
	},
}

// GetAlibabaAlihouseExistinghomeHouseStoreentrustQueryResult() 从对象池中获取AlibabaAlihouseExistinghomeHouseStoreentrustQueryResult
func GetAlibabaAlihouseExistinghomeHouseStoreentrustQueryResult() *AlibabaAlihouseExistinghomeHouseStoreentrustQueryResult {
	return poolAlibabaAlihouseExistinghomeHouseStoreentrustQueryResult.Get().(*AlibabaAlihouseExistinghomeHouseStoreentrustQueryResult)
}

// ReleaseAlibabaAlihouseExistinghomeHouseStoreentrustQueryResult 释放AlibabaAlihouseExistinghomeHouseStoreentrustQueryResult
func ReleaseAlibabaAlihouseExistinghomeHouseStoreentrustQueryResult(v *AlibabaAlihouseExistinghomeHouseStoreentrustQueryResult) {
	v.Code = ""
	v.Msg = ""
	v.Data = nil
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeHouseStoreentrustQueryResult.Put(v)
}
