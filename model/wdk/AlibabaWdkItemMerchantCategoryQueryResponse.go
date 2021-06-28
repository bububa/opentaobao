package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询商品的商家叶子类目 APIResponse
alibaba.wdk.item.merchant.category.query

查询商品的商家叶子类目
*/
type AlibabaWdkItemMerchantCategoryQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkItemMerchantCategoryQueryResponse `json:"alibaba_wdk_item_merchant_category_query_response,omitempty"` 
    AlibabaWdkItemMerchantCategoryQueryResponse
}

/* model for simplify = false
type AlibabaWdkItemMerchantCategoryQueryResponse struct {

    // 结果
    
    Result  *struct {
        WdkOpenSkuMerchantCatServiceQueryResult  *WdkOpenSkuMerchantCatServiceQueryResult `json:"wdk_open_sku_merchant_cat_service_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkItemMerchantCategoryQueryResponse struct {

    // 结果
    Result   *WdkOpenSkuMerchantCatServiceQueryResult `json:"result,omitempty"`

}
