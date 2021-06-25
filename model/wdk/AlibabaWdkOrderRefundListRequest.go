package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口交易退款批量查询 APIRequest
alibaba.wdk.order.refund.list

按照条件查询退款数据
*/
type AlibabaWdkOrderRefundListRequest struct {
    model.Params

    // 查询条件
    batchQueryRefundRequest   *BatchQueryRefundRequest 

}

func NewAlibabaWdkOrderRefundListRequest() *AlibabaWdkOrderRefundListRequest{
    return &AlibabaWdkOrderRefundListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkOrderRefundListRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.refund.list"
}

func (r AlibabaWdkOrderRefundListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkOrderRefundListRequest) SetBatchQueryRefundRequest(batchQueryRefundRequest *BatchQueryRefundRequest) error {
    r.batchQueryRefundRequest = batchQueryRefundRequest
    r.Set("batch_query_refund_request", batchQueryRefundRequest)
    return nil
}

func (r AlibabaWdkOrderRefundListRequest) GetBatchQueryRefundRequest() *BatchQueryRefundRequest {
    return r.batchQueryRefundRequest
}

