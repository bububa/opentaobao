package wdk

import (
	"sync"
)

// AlibabaWdkSkuBarcodeQueryApiResults 结构体
type AlibabaWdkSkuBarcodeQueryApiResults struct {
	// 条码结果集合
	Models []BarcodeBo `json:"models,omitempty" xml:"models>barcode_bo,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功，根据该字段判断是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuBarcodeQueryApiResults = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuBarcodeQueryApiResults)
	},
}

// GetAlibabaWdkSkuBarcodeQueryApiResults() 从对象池中获取AlibabaWdkSkuBarcodeQueryApiResults
func GetAlibabaWdkSkuBarcodeQueryApiResults() *AlibabaWdkSkuBarcodeQueryApiResults {
	return poolAlibabaWdkSkuBarcodeQueryApiResults.Get().(*AlibabaWdkSkuBarcodeQueryApiResults)
}

// ReleaseAlibabaWdkSkuBarcodeQueryApiResults 释放AlibabaWdkSkuBarcodeQueryApiResults
func ReleaseAlibabaWdkSkuBarcodeQueryApiResults(v *AlibabaWdkSkuBarcodeQueryApiResults) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolAlibabaWdkSkuBarcodeQueryApiResults.Put(v)
}
