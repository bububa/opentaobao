package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店查询接口 APIResponse
alibaba.wdk.shop.query

根据门店code查询门店信息
*/
type AlibabaWdkShopQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkShopQueryResponse `json:"alibaba_wdk_shop_query_response,omitempty"` 
    AlibabaWdkShopQueryResponse
}

/* model for simplify = false
type AlibabaWdkShopQueryResponse struct {

    // 调用结果
    
    Result  *struct {
        AlibabaWdkShopQueryApiResults  *AlibabaWdkShopQueryApiResults `json:"alibaba_wdk_shop_query_api_results,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkShopQueryResponse struct {

    // 调用结果
    Result   *AlibabaWdkShopQueryApiResults `json:"result,omitempty"`

}
