package wdk

import (
	"sync"
)

// WdkOpenSkuMerchantCatServiceQueryResult 结构体
type WdkOpenSkuMerchantCatServiceQueryResult struct {
	// 结果描述
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 结果码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// {&#34;skuCode&#34;:&#34;categoryCode&#34;}，skuCode和商家叶子类目编码映射关系
	SkuCodeCategoryCodeMap string `json:"sku_code_category_code_map,omitempty" xml:"sku_code_category_code_map,omitempty"`
	// 成功或失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolWdkOpenSkuMerchantCatServiceQueryResult = sync.Pool{
	New: func() any {
		return new(WdkOpenSkuMerchantCatServiceQueryResult)
	},
}

// GetWdkOpenSkuMerchantCatServiceQueryResult() 从对象池中获取WdkOpenSkuMerchantCatServiceQueryResult
func GetWdkOpenSkuMerchantCatServiceQueryResult() *WdkOpenSkuMerchantCatServiceQueryResult {
	return poolWdkOpenSkuMerchantCatServiceQueryResult.Get().(*WdkOpenSkuMerchantCatServiceQueryResult)
}

// ReleaseWdkOpenSkuMerchantCatServiceQueryResult 释放WdkOpenSkuMerchantCatServiceQueryResult
func ReleaseWdkOpenSkuMerchantCatServiceQueryResult(v *WdkOpenSkuMerchantCatServiceQueryResult) {
	v.ReturnMsg = ""
	v.ReturnCode = ""
	v.SkuCodeCategoryCodeMap = ""
	v.Success = false
	poolWdkOpenSkuMerchantCatServiceQueryResult.Put(v)
}
