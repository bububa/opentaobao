package tmallcarenter

import (
	"sync"
)

// TmallCarcenterVehicleinfoRegisterResult 结构体
type TmallCarcenterVehicleinfoRegisterResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// gmtCurrentTime
	GmtCurrentTime int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
	// costTime
	CostTime int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallCarcenterVehicleinfoRegisterResult = sync.Pool{
	New: func() any {
		return new(TmallCarcenterVehicleinfoRegisterResult)
	},
}

// GetTmallCarcenterVehicleinfoRegisterResult() 从对象池中获取TmallCarcenterVehicleinfoRegisterResult
func GetTmallCarcenterVehicleinfoRegisterResult() *TmallCarcenterVehicleinfoRegisterResult {
	return poolTmallCarcenterVehicleinfoRegisterResult.Get().(*TmallCarcenterVehicleinfoRegisterResult)
}

// ReleaseTmallCarcenterVehicleinfoRegisterResult 释放TmallCarcenterVehicleinfoRegisterResult
func ReleaseTmallCarcenterVehicleinfoRegisterResult(v *TmallCarcenterVehicleinfoRegisterResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.GmtCurrentTime = 0
	v.CostTime = 0
	v.Object = false
	v.Success = false
	poolTmallCarcenterVehicleinfoRegisterResult.Put(v)
}
