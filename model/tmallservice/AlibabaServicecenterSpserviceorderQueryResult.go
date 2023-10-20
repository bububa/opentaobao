package tmallservice

import (
	"sync"
)

// AlibabaServicecenterSpserviceorderQueryResult 结构体
type AlibabaServicecenterSpserviceorderQueryResult struct {
	// 错误描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 分页数据
	ResultData *Paged `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaServicecenterSpserviceorderQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterSpserviceorderQueryResult)
	},
}

// GetAlibabaServicecenterSpserviceorderQueryResult() 从对象池中获取AlibabaServicecenterSpserviceorderQueryResult
func GetAlibabaServicecenterSpserviceorderQueryResult() *AlibabaServicecenterSpserviceorderQueryResult {
	return poolAlibabaServicecenterSpserviceorderQueryResult.Get().(*AlibabaServicecenterSpserviceorderQueryResult)
}

// ReleaseAlibabaServicecenterSpserviceorderQueryResult 释放AlibabaServicecenterSpserviceorderQueryResult
func ReleaseAlibabaServicecenterSpserviceorderQueryResult(v *AlibabaServicecenterSpserviceorderQueryResult) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.ResultData = nil
	v.Success = false
	poolAlibabaServicecenterSpserviceorderQueryResult.Put(v)
}
