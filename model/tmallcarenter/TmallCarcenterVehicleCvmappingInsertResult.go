package tmallcarenter

import (
	"sync"
)

// TmallCarcenterVehicleCvmappingInsertResult 结构体
type TmallCarcenterVehicleCvmappingInsertResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
	// gmtCurrentTime
	GmtCurrentTime int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
	// costTime
	CostTime int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallCarcenterVehicleCvmappingInsertResult = sync.Pool{
	New: func() any {
		return new(TmallCarcenterVehicleCvmappingInsertResult)
	},
}

// GetTmallCarcenterVehicleCvmappingInsertResult() 从对象池中获取TmallCarcenterVehicleCvmappingInsertResult
func GetTmallCarcenterVehicleCvmappingInsertResult() *TmallCarcenterVehicleCvmappingInsertResult {
	return poolTmallCarcenterVehicleCvmappingInsertResult.Get().(*TmallCarcenterVehicleCvmappingInsertResult)
}

// ReleaseTmallCarcenterVehicleCvmappingInsertResult 释放TmallCarcenterVehicleCvmappingInsertResult
func ReleaseTmallCarcenterVehicleCvmappingInsertResult(v *TmallCarcenterVehicleCvmappingInsertResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Object = ""
	v.GmtCurrentTime = 0
	v.CostTime = 0
	v.Success = false
	poolTmallCarcenterVehicleCvmappingInsertResult.Put(v)
}
