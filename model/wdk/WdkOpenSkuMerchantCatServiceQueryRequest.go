package wdk

// WdkOpenSkuMerchantCatServiceQueryRequest 
/* model for simplify = false
type WdkOpenSkuMerchantCatServiceQueryRequest struct {

    // 商品skuCode
    
    SkuCodeList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"sku_code_list,omitempty"`
    

}
*/

// WdkOpenSkuMerchantCatServiceQueryRequest 
type WdkOpenSkuMerchantCatServiceQueryRequest struct {

    // 商品skuCode
    SkuCodeList   []string `json:"sku_code_list,omitempty"`

}
