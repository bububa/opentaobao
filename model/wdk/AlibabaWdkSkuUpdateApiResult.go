package wdk

import (
	"sync"
)

// AlibabaWdkSkuUpdateApiResult 结构体
type AlibabaWdkSkuUpdateApiResult struct {
	// sku编码
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// sku商品操作错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// sku商品操作错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 聚合之后的产品id，商家需要保存该值
	ProductId string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// sku商品操作成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaWdkSkuUpdateApiResult = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuUpdateApiResult)
	},
}

// GetAlibabaWdkSkuUpdateApiResult() 从对象池中获取AlibabaWdkSkuUpdateApiResult
func GetAlibabaWdkSkuUpdateApiResult() *AlibabaWdkSkuUpdateApiResult {
	return poolAlibabaWdkSkuUpdateApiResult.Get().(*AlibabaWdkSkuUpdateApiResult)
}

// ReleaseAlibabaWdkSkuUpdateApiResult 释放AlibabaWdkSkuUpdateApiResult
func ReleaseAlibabaWdkSkuUpdateApiResult(v *AlibabaWdkSkuUpdateApiResult) {
	v.Model = ""
	v.ErrMsg = ""
	v.ErrCode = ""
	v.ProductId = ""
	v.Success = false
	poolAlibabaWdkSkuUpdateApiResult.Put(v)
}
