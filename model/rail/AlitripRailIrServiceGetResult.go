package rail

import (
	"sync"
)

// AlitripRailIrServiceGetResult 结构体
type AlitripRailIrServiceGetResult struct {
	// 仓位坐席数组
	Modules []Modules `json:"modules,omitempty" xml:"modules>modules,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlitripRailIrServiceGetResult = sync.Pool{
	New: func() any {
		return new(AlitripRailIrServiceGetResult)
	},
}

// GetAlitripRailIrServiceGetResult() 从对象池中获取AlitripRailIrServiceGetResult
func GetAlitripRailIrServiceGetResult() *AlitripRailIrServiceGetResult {
	return poolAlitripRailIrServiceGetResult.Get().(*AlitripRailIrServiceGetResult)
}

// ReleaseAlitripRailIrServiceGetResult 释放AlitripRailIrServiceGetResult
func ReleaseAlitripRailIrServiceGetResult(v *AlitripRailIrServiceGetResult) {
	v.Modules = v.Modules[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlitripRailIrServiceGetResult.Put(v)
}
