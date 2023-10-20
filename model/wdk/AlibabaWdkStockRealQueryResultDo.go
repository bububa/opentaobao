package wdk

import (
	"sync"
)

// AlibabaWdkStockRealQueryResultDo 结构体
type AlibabaWdkStockRealQueryResultDo struct {
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 错误码
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// result
	Data *InventoryTopResultBo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkStockRealQueryResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaWdkStockRealQueryResultDo)
	},
}

// GetAlibabaWdkStockRealQueryResultDo() 从对象池中获取AlibabaWdkStockRealQueryResultDo
func GetAlibabaWdkStockRealQueryResultDo() *AlibabaWdkStockRealQueryResultDo {
	return poolAlibabaWdkStockRealQueryResultDo.Get().(*AlibabaWdkStockRealQueryResultDo)
}

// ReleaseAlibabaWdkStockRealQueryResultDo 释放AlibabaWdkStockRealQueryResultDo
func ReleaseAlibabaWdkStockRealQueryResultDo(v *AlibabaWdkStockRealQueryResultDo) {
	v.ErrMsg = ""
	v.ErrCode = 0
	v.Data = nil
	v.Success = false
	poolAlibabaWdkStockRealQueryResultDo.Put(v)
}
