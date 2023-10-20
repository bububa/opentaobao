package alsc

import (
	"sync"
)

// AlibabaAlscSaasCodecCodeAttrsQueryResult 结构体
type AlibabaAlscSaasCodecCodeAttrsQueryResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回素材id
	Data *CodeBizAttributeDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlscSaasCodecCodeAttrsQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlscSaasCodecCodeAttrsQueryResult)
	},
}

// GetAlibabaAlscSaasCodecCodeAttrsQueryResult() 从对象池中获取AlibabaAlscSaasCodecCodeAttrsQueryResult
func GetAlibabaAlscSaasCodecCodeAttrsQueryResult() *AlibabaAlscSaasCodecCodeAttrsQueryResult {
	return poolAlibabaAlscSaasCodecCodeAttrsQueryResult.Get().(*AlibabaAlscSaasCodecCodeAttrsQueryResult)
}

// ReleaseAlibabaAlscSaasCodecCodeAttrsQueryResult 释放AlibabaAlscSaasCodecCodeAttrsQueryResult
func ReleaseAlibabaAlscSaasCodecCodeAttrsQueryResult(v *AlibabaAlscSaasCodecCodeAttrsQueryResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Data = nil
	v.Success = false
	poolAlibabaAlscSaasCodecCodeAttrsQueryResult.Put(v)
}
