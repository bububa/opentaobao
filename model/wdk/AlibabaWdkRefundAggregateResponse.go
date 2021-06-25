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
    Response *AlibabaWdkRefundAggregateResponse `json:"alibaba_wdk_refund_aggregate_response,omitempty"`
}

type AlibabaWdkRefundAggregateResponse struct {

    // result
    Result   *RefundAggregateQueryResult `json:"result,omitempty"`

}
