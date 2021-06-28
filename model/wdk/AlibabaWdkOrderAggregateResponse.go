package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达订单按门店机台号聚合查询 APIResponse
alibaba.wdk.order.aggregate

淘鲜达订单按门店机台号聚合查询
*/
type AlibabaWdkOrderAggregateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkOrderAggregateResponse `json:"alibaba_wdk_order_aggregate_response,omitempty"` 
    AlibabaWdkOrderAggregateResponse
}

/* model for simplify = false
type AlibabaWdkOrderAggregateResponse struct {

    // result
    
    Result  *struct {
        OrderAggregateQueryResult  *OrderAggregateQueryResult `json:"order_aggregate_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkOrderAggregateResponse struct {

    // result
    Result   *OrderAggregateQueryResult `json:"result,omitempty"`

}
