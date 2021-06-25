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
    Response *AlibabaWdkItemMerchantCategoryQueryResponse `json:"alibaba_wdk_item_merchant_category_query_response,omitempty"`
}

type AlibabaWdkItemMerchantCategoryQueryResponse struct {

    // 结果
    Result   *WdkOpenSkuMerchantCatServiceQueryResult `json:"result,omitempty"`

}
