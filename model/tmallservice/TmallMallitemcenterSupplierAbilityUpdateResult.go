package tmallservice

import (
	"sync"
)

// TmallMallitemcenterSupplierAbilityUpdateResult 结构体
type TmallMallitemcenterSupplierAbilityUpdateResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误类型
	ErrorType int64 `json:"error_type,omitempty" xml:"error_type,omitempty"`
	// 结果
	ResultData bool `json:"result_data,omitempty" xml:"result_data,omitempty"`
	// true或false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallMallitemcenterSupplierAbilityUpdateResult = sync.Pool{
	New: func() any {
		return new(TmallMallitemcenterSupplierAbilityUpdateResult)
	},
}

// GetTmallMallitemcenterSupplierAbilityUpdateResult() 从对象池中获取TmallMallitemcenterSupplierAbilityUpdateResult
func GetTmallMallitemcenterSupplierAbilityUpdateResult() *TmallMallitemcenterSupplierAbilityUpdateResult {
	return poolTmallMallitemcenterSupplierAbilityUpdateResult.Get().(*TmallMallitemcenterSupplierAbilityUpdateResult)
}

// ReleaseTmallMallitemcenterSupplierAbilityUpdateResult 释放TmallMallitemcenterSupplierAbilityUpdateResult
func ReleaseTmallMallitemcenterSupplierAbilityUpdateResult(v *TmallMallitemcenterSupplierAbilityUpdateResult) {
	v.Message = ""
	v.ErrorCode = ""
	v.ErrorType = 0
	v.ResultData = false
	v.Success = false
	poolTmallMallitemcenterSupplierAbilityUpdateResult.Put(v)
}
