package wdk

import (
	"sync"
)

// WdkOpenSkuMerchantCatServiceQueryRequest 结构体
type WdkOpenSkuMerchantCatServiceQueryRequest struct {
	// 商品skuCode
	SkuCodeList []string `json:"sku_code_list,omitempty" xml:"sku_code_list>string,omitempty"`
}

var poolWdkOpenSkuMerchantCatServiceQueryRequest = sync.Pool{
	New: func() any {
		return new(WdkOpenSkuMerchantCatServiceQueryRequest)
	},
}

// GetWdkOpenSkuMerchantCatServiceQueryRequest() 从对象池中获取WdkOpenSkuMerchantCatServiceQueryRequest
func GetWdkOpenSkuMerchantCatServiceQueryRequest() *WdkOpenSkuMerchantCatServiceQueryRequest {
	return poolWdkOpenSkuMerchantCatServiceQueryRequest.Get().(*WdkOpenSkuMerchantCatServiceQueryRequest)
}

// ReleaseWdkOpenSkuMerchantCatServiceQueryRequest 释放WdkOpenSkuMerchantCatServiceQueryRequest
func ReleaseWdkOpenSkuMerchantCatServiceQueryRequest(v *WdkOpenSkuMerchantCatServiceQueryRequest) {
	v.SkuCodeList = v.SkuCodeList[:0]
	poolWdkOpenSkuMerchantCatServiceQueryRequest.Put(v)
}
