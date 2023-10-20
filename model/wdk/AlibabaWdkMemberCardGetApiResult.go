package wdk

import (
	"sync"
)

// AlibabaWdkMemberCardGetApiResult 结构体
type AlibabaWdkMemberCardGetApiResult struct {
	// 错误消息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误结果码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 会员信息模型
	Model *MemberInfoDo `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkMemberCardGetApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMemberCardGetApiResult)
	},
}

// GetAlibabaWdkMemberCardGetApiResult() 从对象池中获取AlibabaWdkMemberCardGetApiResult
func GetAlibabaWdkMemberCardGetApiResult() *AlibabaWdkMemberCardGetApiResult {
	return poolAlibabaWdkMemberCardGetApiResult.Get().(*AlibabaWdkMemberCardGetApiResult)
}

// ReleaseAlibabaWdkMemberCardGetApiResult 释放AlibabaWdkMemberCardGetApiResult
func ReleaseAlibabaWdkMemberCardGetApiResult(v *AlibabaWdkMemberCardGetApiResult) {
	v.ErrMsg = ""
	v.ErrCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkMemberCardGetApiResult.Put(v)
}
