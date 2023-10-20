package ascpffo

import (
	"sync"
)

// AliexpressAscpPoQueryResult 结构体
type AliexpressAscpPoQueryResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 采购单DTO
	Data *ErpPurchaseOrderDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAliexpressAscpPoQueryResult = sync.Pool{
	New: func() any {
		return new(AliexpressAscpPoQueryResult)
	},
}

// GetAliexpressAscpPoQueryResult() 从对象池中获取AliexpressAscpPoQueryResult
func GetAliexpressAscpPoQueryResult() *AliexpressAscpPoQueryResult {
	return poolAliexpressAscpPoQueryResult.Get().(*AliexpressAscpPoQueryResult)
}

// ReleaseAliexpressAscpPoQueryResult 释放AliexpressAscpPoQueryResult
func ReleaseAliexpressAscpPoQueryResult(v *AliexpressAscpPoQueryResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Data = nil
	v.Success = false
	poolAliexpressAscpPoQueryResult.Put(v)
}
