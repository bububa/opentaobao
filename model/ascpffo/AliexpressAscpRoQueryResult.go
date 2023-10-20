package ascpffo

import (
	"sync"
)

// AliexpressAscpRoQueryResult 结构体
type AliexpressAscpRoQueryResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// DTO
	Data *ErpReturnOrderDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAliexpressAscpRoQueryResult = sync.Pool{
	New: func() any {
		return new(AliexpressAscpRoQueryResult)
	},
}

// GetAliexpressAscpRoQueryResult() 从对象池中获取AliexpressAscpRoQueryResult
func GetAliexpressAscpRoQueryResult() *AliexpressAscpRoQueryResult {
	return poolAliexpressAscpRoQueryResult.Get().(*AliexpressAscpRoQueryResult)
}

// ReleaseAliexpressAscpRoQueryResult 释放AliexpressAscpRoQueryResult
func ReleaseAliexpressAscpRoQueryResult(v *AliexpressAscpRoQueryResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolAliexpressAscpRoQueryResult.Put(v)
}
