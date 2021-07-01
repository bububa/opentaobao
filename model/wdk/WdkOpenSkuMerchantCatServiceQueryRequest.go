package wdk

// WdkOpenSkuMerchantCatServiceQueryRequest 结构体
type WdkOpenSkuMerchantCatServiceQueryRequest struct {
	// 商品skuCode
	SkuCodeList []string `json:"sku_code_list,omitempty" xml:"sku_code_list>string,omitempty"`
}
