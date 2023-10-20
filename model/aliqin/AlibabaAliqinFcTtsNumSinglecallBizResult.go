package aliqin

import (
	"sync"
)

// AlibabaAliqinFcTtsNumSinglecallBizResult 结构体
type AlibabaAliqinFcTtsNumSinglecallBizResult struct {
	// 返回结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回信息描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// true表示成功，false表示失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAliqinFcTtsNumSinglecallBizResult = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcTtsNumSinglecallBizResult)
	},
}

// GetAlibabaAliqinFcTtsNumSinglecallBizResult() 从对象池中获取AlibabaAliqinFcTtsNumSinglecallBizResult
func GetAlibabaAliqinFcTtsNumSinglecallBizResult() *AlibabaAliqinFcTtsNumSinglecallBizResult {
	return poolAlibabaAliqinFcTtsNumSinglecallBizResult.Get().(*AlibabaAliqinFcTtsNumSinglecallBizResult)
}

// ReleaseAlibabaAliqinFcTtsNumSinglecallBizResult 释放AlibabaAliqinFcTtsNumSinglecallBizResult
func ReleaseAlibabaAliqinFcTtsNumSinglecallBizResult(v *AlibabaAliqinFcTtsNumSinglecallBizResult) {
	v.Model = ""
	v.Msg = ""
	v.ErrCode = ""
	v.Success = false
	poolAlibabaAliqinFcTtsNumSinglecallBizResult.Put(v)
}
