package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询商品 APIResponse
alibaba.wdk.sku.query

查询商品
*/
type AlibabaWdkSkuQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSkuQueryResponse `json:"alibaba_wdk_sku_query_response,omitempty"` 
    AlibabaWdkSkuQueryResponse
}

/* model for simplify = false
type AlibabaWdkSkuQueryResponse struct {

    // 调用结果
    
    Result  *struct {
        AlibabaWdkSkuQueryApiResults  *AlibabaWdkSkuQueryApiResults `json:"alibaba_wdk_sku_query_api_results,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkSkuQueryResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuQueryApiResults `json:"result,omitempty"`

}
