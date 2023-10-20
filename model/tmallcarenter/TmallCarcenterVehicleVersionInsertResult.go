package tmallcarenter

import (
	"sync"
)

// TmallCarcenterVehicleVersionInsertResult 结构体
type TmallCarcenterVehicleVersionInsertResult struct {
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

var poolTmallCarcenterVehicleVersionInsertResult = sync.Pool{
	New: func() any {
		return new(TmallCarcenterVehicleVersionInsertResult)
	},
}

// GetTmallCarcenterVehicleVersionInsertResult() 从对象池中获取TmallCarcenterVehicleVersionInsertResult
func GetTmallCarcenterVehicleVersionInsertResult() *TmallCarcenterVehicleVersionInsertResult {
	return poolTmallCarcenterVehicleVersionInsertResult.Get().(*TmallCarcenterVehicleVersionInsertResult)
}

// ReleaseTmallCarcenterVehicleVersionInsertResult 释放TmallCarcenterVehicleVersionInsertResult
func ReleaseTmallCarcenterVehicleVersionInsertResult(v *TmallCarcenterVehicleVersionInsertResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Object = ""
	v.GmtCurrentTime = 0
	v.CostTime = 0
	v.Success = false
	poolTmallCarcenterVehicleVersionInsertResult.Put(v)
}
