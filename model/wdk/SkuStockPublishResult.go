package wdk

import (
	"sync"
)

// SkuStockPublishResult 结构体
type SkuStockPublishResult struct {
	// 具体的错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 入参中对应的幂等单据号
	BillNo string `json:"bill_no,omitempty" xml:"bill_no,omitempty"`
	// 具体的更新失败原因
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// bill_no对应的操作结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolSkuStockPublishResult = sync.Pool{
	New: func() any {
		return new(SkuStockPublishResult)
	},
}

// GetSkuStockPublishResult() 从对象池中获取SkuStockPublishResult
func GetSkuStockPublishResult() *SkuStockPublishResult {
	return poolSkuStockPublishResult.Get().(*SkuStockPublishResult)
}

// ReleaseSkuStockPublishResult 释放SkuStockPublishResult
func ReleaseSkuStockPublishResult(v *SkuStockPublishResult) {
	v.ErrorCode = ""
	v.BillNo = ""
	v.ErrMsg = ""
	v.Result = false
	poolSkuStockPublishResult.Put(v)
}
