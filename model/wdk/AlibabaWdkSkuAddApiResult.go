package wdk

import (
	"sync"
)

// AlibabaWdkSkuAddApiResult 结构体
type AlibabaWdkSkuAddApiResult struct {
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// skuCode商品编码
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 聚合之后的产品id，商家需要保存该值
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 单个商品新增是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuAddApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuAddApiResult)
	},
}

// GetAlibabaWdkSkuAddApiResult() 从对象池中获取AlibabaWdkSkuAddApiResult
func GetAlibabaWdkSkuAddApiResult() *AlibabaWdkSkuAddApiResult {
	return poolAlibabaWdkSkuAddApiResult.Get().(*AlibabaWdkSkuAddApiResult)
}

// ReleaseAlibabaWdkSkuAddApiResult 释放AlibabaWdkSkuAddApiResult
func ReleaseAlibabaWdkSkuAddApiResult(v *AlibabaWdkSkuAddApiResult) {
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Model = ""
	v.ProductId = ""
	v.Success = false
	poolAlibabaWdkSkuAddApiResult.Put(v)
}
