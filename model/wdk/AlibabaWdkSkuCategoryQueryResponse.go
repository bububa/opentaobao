package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家类目获取接口 APIResponse
alibaba.wdk.sku.category.query

商家类目获取接口
*/
type AlibabaWdkSkuCategoryQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSkuCategoryQueryResponse `json:"alibaba_wdk_sku_category_query_response,omitempty"` 
    AlibabaWdkSkuCategoryQueryResponse
}

/* model for simplify = false
type AlibabaWdkSkuCategoryQueryResponse struct {

    // 调用结果
    
    Result  *struct {
        AlibabaWdkSkuCategoryQueryApiResult  *AlibabaWdkSkuCategoryQueryApiResult `json:"alibaba_wdk_sku_category_query_api_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkSkuCategoryQueryResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuCategoryQueryApiResult `json:"result,omitempty"`

}
