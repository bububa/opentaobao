package tmallservice

import (
	"sync"
)

// TmallSscSupplyplatformCapacityEditResult 结构体
type TmallSscSupplyplatformCapacityEditResult struct {
	// 对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallSscSupplyplatformCapacityEditResult = sync.Pool{
	New: func() any {
		return new(TmallSscSupplyplatformCapacityEditResult)
	},
}

// GetTmallSscSupplyplatformCapacityEditResult() 从对象池中获取TmallSscSupplyplatformCapacityEditResult
func GetTmallSscSupplyplatformCapacityEditResult() *TmallSscSupplyplatformCapacityEditResult {
	return poolTmallSscSupplyplatformCapacityEditResult.Get().(*TmallSscSupplyplatformCapacityEditResult)
}

// ReleaseTmallSscSupplyplatformCapacityEditResult 释放TmallSscSupplyplatformCapacityEditResult
func ReleaseTmallSscSupplyplatformCapacityEditResult(v *TmallSscSupplyplatformCapacityEditResult) {
	v.DisplayMsg = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = ""
	poolTmallSscSupplyplatformCapacityEditResult.Put(v)
}
