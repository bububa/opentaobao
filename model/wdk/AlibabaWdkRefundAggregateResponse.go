package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达退款单按门店聚合查询 APIResponse
alibaba.wdk.refund.aggregate

淘鲜达退款单按门店聚合查询
*/
type AlibabaWdkRefundAggregateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkRefundAggregateResponse `json:"alibaba_wdk_refund_aggregate_response,omitempty"` 
    AlibabaWdkRefundAggregateResponse
}

/* model for simplify = false
type AlibabaWdkRefundAggregateResponse struct {

    // result
    
    Result  *struct {
        RefundAggregateQueryResult  *RefundAggregateQueryResult `json:"refund_aggregate_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkRefundAggregateResponse struct {

    // result
    Result   *RefundAggregateQueryResult `json:"result,omitempty"`

}
