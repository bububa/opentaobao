package alihouse

import (
	"sync"
)

// AlibabaAlihouseExistinghomeCompanyinfoSubmitResult 结构体
type AlibabaAlihouseExistinghomeCompanyinfoSubmitResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 进件id
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolAlibabaAlihouseExistinghomeCompanyinfoSubmitResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseExistinghomeCompanyinfoSubmitResult)
	},
}

// GetAlibabaAlihouseExistinghomeCompanyinfoSubmitResult() 从对象池中获取AlibabaAlihouseExistinghomeCompanyinfoSubmitResult
func GetAlibabaAlihouseExistinghomeCompanyinfoSubmitResult() *AlibabaAlihouseExistinghomeCompanyinfoSubmitResult {
	return poolAlibabaAlihouseExistinghomeCompanyinfoSubmitResult.Get().(*AlibabaAlihouseExistinghomeCompanyinfoSubmitResult)
}

// ReleaseAlibabaAlihouseExistinghomeCompanyinfoSubmitResult 释放AlibabaAlihouseExistinghomeCompanyinfoSubmitResult
func ReleaseAlibabaAlihouseExistinghomeCompanyinfoSubmitResult(v *AlibabaAlihouseExistinghomeCompanyinfoSubmitResult) {
	v.Code = ""
	v.Msg = ""
	v.Data = 0
	v.IsSuccess = false
	poolAlibabaAlihouseExistinghomeCompanyinfoSubmitResult.Put(v)
}
