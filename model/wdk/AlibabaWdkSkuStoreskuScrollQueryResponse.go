package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
门店商品批量查询接口 APIResponse
alibaba.wdk.sku.storesku.scroll.query

提供门店商品批量查询接口
*/
type AlibabaWdkSkuStoreskuScrollQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSkuStoreskuScrollQueryResponse `json:"alibaba_wdk_sku_storesku_scroll_query_response,omitempty"` 
    AlibabaWdkSkuStoreskuScrollQueryResponse
}

/* model for simplify = false
type AlibabaWdkSkuStoreskuScrollQueryResponse struct {

    // 请求结果
    
    Result  *struct {
        ApiScrollPageResult  *ApiScrollPageResult `json:"api_scroll_page_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkSkuStoreskuScrollQueryResponse struct {

    // 请求结果
    Result   *ApiScrollPageResult `json:"result,omitempty"`

}
