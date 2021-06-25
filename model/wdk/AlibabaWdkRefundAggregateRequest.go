package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达退款单按门店聚合查询 APIRequest
alibaba.wdk.refund.aggregate

淘鲜达退款单按门店聚合查询
*/
type AlibabaWdkRefundAggregateRequest struct {
    model.Params

    // 系统自动生成
    refundAggregateQueryRequest   *RefundAggregateQueryRequest 

}

func NewAlibabaWdkRefundAggregateRequest() *AlibabaWdkRefundAggregateRequest{
    return &AlibabaWdkRefundAggregateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkRefundAggregateRequest) GetApiMethodName() string {
    return "alibaba.wdk.refund.aggregate"
}

func (r AlibabaWdkRefundAggregateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkRefundAggregateRequest) SetRefundAggregateQueryRequest(refundAggregateQueryRequest *RefundAggregateQueryRequest) error {
    r.refundAggregateQueryRequest = refundAggregateQueryRequest
    r.Set("refund_aggregate_query_request", refundAggregateQueryRequest)
    return nil
}

func (r AlibabaWdkRefundAggregateRequest) GetRefundAggregateQueryRequest() *RefundAggregateQueryRequest {
    return r.refundAggregateQueryRequest
}

