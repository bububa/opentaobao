package alihouse

import (
	"sync"
)

// AlibabaAlihouseStorePunishInfoQueryResult 结构体
type AlibabaAlihouseStorePunishInfoQueryResult struct {
	// dto
	Data []StorePunishDto `json:"data,omitempty" xml:"data>store_punish_dto,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// message
	ResultMessage string `json:"result_message,omitempty" xml:"result_message,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

var poolAlibabaAlihouseStorePunishInfoQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseStorePunishInfoQueryResult)
	},
}

// GetAlibabaAlihouseStorePunishInfoQueryResult() 从对象池中获取AlibabaAlihouseStorePunishInfoQueryResult
func GetAlibabaAlihouseStorePunishInfoQueryResult() *AlibabaAlihouseStorePunishInfoQueryResult {
	return poolAlibabaAlihouseStorePunishInfoQueryResult.Get().(*AlibabaAlihouseStorePunishInfoQueryResult)
}

// ReleaseAlibabaAlihouseStorePunishInfoQueryResult 释放AlibabaAlihouseStorePunishInfoQueryResult
func ReleaseAlibabaAlihouseStorePunishInfoQueryResult(v *AlibabaAlihouseStorePunishInfoQueryResult) {
	v.Data = v.Data[:0]
	v.ResultCode = ""
	v.ResultMessage = ""
	v.ResultSuccess = false
	poolAlibabaAlihouseStorePunishInfoQueryResult.Put(v)
}
