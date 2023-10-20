package wdk

import (
	"sync"
)

// AlibabaWdkMemberQrcodeIdentifyMtopResult 结构体
type AlibabaWdkMemberQrcodeIdentifyMtopResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *MemberInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkMemberQrcodeIdentifyMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMemberQrcodeIdentifyMtopResult)
	},
}

// GetAlibabaWdkMemberQrcodeIdentifyMtopResult() 从对象池中获取AlibabaWdkMemberQrcodeIdentifyMtopResult
func GetAlibabaWdkMemberQrcodeIdentifyMtopResult() *AlibabaWdkMemberQrcodeIdentifyMtopResult {
	return poolAlibabaWdkMemberQrcodeIdentifyMtopResult.Get().(*AlibabaWdkMemberQrcodeIdentifyMtopResult)
}

// ReleaseAlibabaWdkMemberQrcodeIdentifyMtopResult 释放AlibabaWdkMemberQrcodeIdentifyMtopResult
func ReleaseAlibabaWdkMemberQrcodeIdentifyMtopResult(v *AlibabaWdkMemberQrcodeIdentifyMtopResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolAlibabaWdkMemberQrcodeIdentifyMtopResult.Put(v)
}
