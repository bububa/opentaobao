package interact

import (
	"sync"
)

// AlibabaInteractUserIsloginMtopResult 结构体
type AlibabaInteractUserIsloginMtopResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaInteractUserIsloginMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaInteractUserIsloginMtopResult)
	},
}

// GetAlibabaInteractUserIsloginMtopResult() 从对象池中获取AlibabaInteractUserIsloginMtopResult
func GetAlibabaInteractUserIsloginMtopResult() *AlibabaInteractUserIsloginMtopResult {
	return poolAlibabaInteractUserIsloginMtopResult.Get().(*AlibabaInteractUserIsloginMtopResult)
}

// ReleaseAlibabaInteractUserIsloginMtopResult 释放AlibabaInteractUserIsloginMtopResult
func ReleaseAlibabaInteractUserIsloginMtopResult(v *AlibabaInteractUserIsloginMtopResult) {
	v.Model = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolAlibabaInteractUserIsloginMtopResult.Put(v)
}
