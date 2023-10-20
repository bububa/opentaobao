package baichuan

import (
	"sync"
)

// AlibabaBaichuanCtgUserRelationResult 结构体
type AlibabaBaichuanCtgUserRelationResult struct {
	// 错误信息
	ErrorDetail string `json:"error_detail,omitempty" xml:"error_detail,omitempty"`
	// 返回结果，数值型，1：代表已绑定达人身份，2代表未绑定达人身份
	Module int64 `json:"module,omitempty" xml:"module,omitempty"`
	// 错误码结构体
	ErrorCode *ErrorCode `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 返回成功与否
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaBaichuanCtgUserRelationResult = sync.Pool{
	New: func() any {
		return new(AlibabaBaichuanCtgUserRelationResult)
	},
}

// GetAlibabaBaichuanCtgUserRelationResult() 从对象池中获取AlibabaBaichuanCtgUserRelationResult
func GetAlibabaBaichuanCtgUserRelationResult() *AlibabaBaichuanCtgUserRelationResult {
	return poolAlibabaBaichuanCtgUserRelationResult.Get().(*AlibabaBaichuanCtgUserRelationResult)
}

// ReleaseAlibabaBaichuanCtgUserRelationResult 释放AlibabaBaichuanCtgUserRelationResult
func ReleaseAlibabaBaichuanCtgUserRelationResult(v *AlibabaBaichuanCtgUserRelationResult) {
	v.ErrorDetail = ""
	v.Module = 0
	v.ErrorCode = nil
	v.Success = false
	poolAlibabaBaichuanCtgUserRelationResult.Put(v)
}
