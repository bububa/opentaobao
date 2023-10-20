package fivee

import (
	"sync"
)

// TaobaoFiveeImportproductGetResult 结构体
type TaobaoFiveeImportproductGetResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data *ImportProduct `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoFiveeImportproductGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoFiveeImportproductGetResult)
	},
}

// GetTaobaoFiveeImportproductGetResult() 从对象池中获取TaobaoFiveeImportproductGetResult
func GetTaobaoFiveeImportproductGetResult() *TaobaoFiveeImportproductGetResult {
	return poolTaobaoFiveeImportproductGetResult.Get().(*TaobaoFiveeImportproductGetResult)
}

// ReleaseTaobaoFiveeImportproductGetResult 释放TaobaoFiveeImportproductGetResult
func ReleaseTaobaoFiveeImportproductGetResult(v *TaobaoFiveeImportproductGetResult) {
	v.Message = ""
	v.Code = 0
	v.Data = nil
	v.Success = false
	poolTaobaoFiveeImportproductGetResult.Put(v)
}
