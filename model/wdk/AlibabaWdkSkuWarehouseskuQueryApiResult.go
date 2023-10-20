package wdk

import (
	"sync"
)

// AlibabaWdkSkuWarehouseskuQueryApiResult 结构体
type AlibabaWdkSkuWarehouseskuQueryApiResult struct {
	// 数据集合
	Models []WarehouseSkuDo `json:"models,omitempty" xml:"models>warehouse_sku_do,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误内容
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuWarehouseskuQueryApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuWarehouseskuQueryApiResult)
	},
}

// GetAlibabaWdkSkuWarehouseskuQueryApiResult() 从对象池中获取AlibabaWdkSkuWarehouseskuQueryApiResult
func GetAlibabaWdkSkuWarehouseskuQueryApiResult() *AlibabaWdkSkuWarehouseskuQueryApiResult {
	return poolAlibabaWdkSkuWarehouseskuQueryApiResult.Get().(*AlibabaWdkSkuWarehouseskuQueryApiResult)
}

// ReleaseAlibabaWdkSkuWarehouseskuQueryApiResult 释放AlibabaWdkSkuWarehouseskuQueryApiResult
func ReleaseAlibabaWdkSkuWarehouseskuQueryApiResult(v *AlibabaWdkSkuWarehouseskuQueryApiResult) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkSkuWarehouseskuQueryApiResult.Put(v)
}
