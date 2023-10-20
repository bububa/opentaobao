package tmallgenie

import (
	"sync"
)

// OpsCommonResult 结构体
type OpsCommonResult struct {
	// 错误信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 错误码，0为成功，其余为失败
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

var poolOpsCommonResult = sync.Pool{
	New: func() any {
		return new(OpsCommonResult)
	},
}

// GetOpsCommonResult() 从对象池中获取OpsCommonResult
func GetOpsCommonResult() *OpsCommonResult {
	return poolOpsCommonResult.Get().(*OpsCommonResult)
}

// ReleaseOpsCommonResult 释放OpsCommonResult
func ReleaseOpsCommonResult(v *OpsCommonResult) {
	v.RetMsg = ""
	v.RetCode = 0
	poolOpsCommonResult.Put(v)
}
