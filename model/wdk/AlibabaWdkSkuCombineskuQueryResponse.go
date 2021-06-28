package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
组合商品查询接口 APIResponse
alibaba.wdk.sku.combinesku.query

查询组合商品接口
*/
type AlibabaWdkSkuCombineskuQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSkuCombineskuQueryResponse `json:"alibaba_wdk_sku_combinesku_query_response,omitempty"` 
    AlibabaWdkSkuCombineskuQueryResponse
}

/* model for simplify = false
type AlibabaWdkSkuCombineskuQueryResponse struct {

    // 调用结果
    
    Result  *struct {
        AlibabaWdkSkuCombineskuQueryApiResults  *AlibabaWdkSkuCombineskuQueryApiResults `json:"alibaba_wdk_sku_combinesku_query_api_results,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkSkuCombineskuQueryResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuCombineskuQueryApiResults `json:"result,omitempty"`

}
