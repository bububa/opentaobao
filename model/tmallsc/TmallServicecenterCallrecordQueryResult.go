package tmallsc

import (
	"sync"
)

// TmallServicecenterCallrecordQueryResult 结构体
type TmallServicecenterCallrecordQueryResult struct {
	// 通话记录信息
	Value []ServiceCallRecordCO `json:"value,omitempty" xml:"value>service_call_record_co,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterCallrecordQueryResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterCallrecordQueryResult)
	},
}

// GetTmallServicecenterCallrecordQueryResult() 从对象池中获取TmallServicecenterCallrecordQueryResult
func GetTmallServicecenterCallrecordQueryResult() *TmallServicecenterCallrecordQueryResult {
	return poolTmallServicecenterCallrecordQueryResult.Get().(*TmallServicecenterCallrecordQueryResult)
}

// ReleaseTmallServicecenterCallrecordQueryResult 释放TmallServicecenterCallrecordQueryResult
func ReleaseTmallServicecenterCallrecordQueryResult(v *TmallServicecenterCallrecordQueryResult) {
	v.Value = v.Value[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTmallServicecenterCallrecordQueryResult.Put(v)
}
