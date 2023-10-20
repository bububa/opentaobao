package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeHouseChangeCommunityResult 结构体
type AlibabaAlihouseExistinghomeHouseChangeCommunityResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// msg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 123
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// true
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeHouseChangeCommunityResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeHouseChangeCommunityResult)
	},
}

// GetAlibabaAlihouseExistinghomeHouseChangeCommunityResult() 从对象池中获取AlibabaAlihouseExistinghomeHouseChangeCommunityResult
func GetAlibabaAlihouseExistinghomeHouseChangeCommunityResult() *AlibabaAlihouseExistinghomeHouseChangeCommunityResult {
	return poolAlibabaAlihouseExistinghomeHouseChangeCommunityResult.Get().(*AlibabaAlihouseExistinghomeHouseChangeCommunityResult)
}

// ReleaseAlibabaAlihouseExistinghomeHouseChangeCommunityResult 释放AlibabaAlihouseExistinghomeHouseChangeCommunityResult
func ReleaseAlibabaAlihouseExistinghomeHouseChangeCommunityResult(v *AlibabaAlihouseExistinghomeHouseChangeCommunityResult) {
	v.Code = ""
	v.Msg = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeHouseChangeCommunityResult.Put(v)
}
