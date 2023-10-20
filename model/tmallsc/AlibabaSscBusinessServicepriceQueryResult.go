package tmallsc

import (
	"sync"
)

// AlibabaSscBusinessServicepriceQueryResult 结构体
type AlibabaSscBusinessServicepriceQueryResult struct {
	// 对用户展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 系统内部错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 系统内部错误信息CODE
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 结果对象
	Value *ServicePriceQueryResponse `json:"value,omitempty" xml:"value,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaSscBusinessServicepriceQueryResult = sync.Pool{
	New: func() any {
		return new(AlibabaSscBusinessServicepriceQueryResult)
	},
}

// GetAlibabaSscBusinessServicepriceQueryResult() 从对象池中获取AlibabaSscBusinessServicepriceQueryResult
func GetAlibabaSscBusinessServicepriceQueryResult() *AlibabaSscBusinessServicepriceQueryResult {
	return poolAlibabaSscBusinessServicepriceQueryResult.Get().(*AlibabaSscBusinessServicepriceQueryResult)
}

// ReleaseAlibabaSscBusinessServicepriceQueryResult 释放AlibabaSscBusinessServicepriceQueryResult
func ReleaseAlibabaSscBusinessServicepriceQueryResult(v *AlibabaSscBusinessServicepriceQueryResult) {
	v.DisplayMsg = ""
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.Value = nil
	v.Success = false
	poolAlibabaSscBusinessServicepriceQueryResult.Put(v)
}
