package tmallservice

import (
	"sync"
)

// AlibabaServicecenterSpserviceorderUpdateResult 结构体
type AlibabaServicecenterSpserviceorderUpdateResult struct {
	// 错误描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误描述
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaServicecenterSpserviceorderUpdateResult = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterSpserviceorderUpdateResult)
	},
}

// GetAlibabaServicecenterSpserviceorderUpdateResult() 从对象池中获取AlibabaServicecenterSpserviceorderUpdateResult
func GetAlibabaServicecenterSpserviceorderUpdateResult() *AlibabaServicecenterSpserviceorderUpdateResult {
	return poolAlibabaServicecenterSpserviceorderUpdateResult.Get().(*AlibabaServicecenterSpserviceorderUpdateResult)
}

// ReleaseAlibabaServicecenterSpserviceorderUpdateResult 释放AlibabaServicecenterSpserviceorderUpdateResult
func ReleaseAlibabaServicecenterSpserviceorderUpdateResult(v *AlibabaServicecenterSpserviceorderUpdateResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.DisplayMsg = ""
	v.Success = false
	poolAlibabaServicecenterSpserviceorderUpdateResult.Put(v)
}
