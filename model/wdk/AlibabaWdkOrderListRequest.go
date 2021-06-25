package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口订单拉取 APIRequest
alibaba.wdk.order.list

五道口交易订单拉取接口
*/
type AlibabaWdkOrderListRequest struct {
    model.Params

    // 查询参数
    batchQueryRequest   *BatchQueryRequest 

}

func NewAlibabaWdkOrderListRequest() *AlibabaWdkOrderListRequest{
    return &AlibabaWdkOrderListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkOrderListRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.list"
}

func (r AlibabaWdkOrderListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkOrderListRequest) SetBatchQueryRequest(batchQueryRequest *BatchQueryRequest) error {
    r.batchQueryRequest = batchQueryRequest
    r.Set("batch_query_request", batchQueryRequest)
    return nil
}

func (r AlibabaWdkOrderListRequest) GetBatchQueryRequest() *BatchQueryRequest {
    return r.batchQueryRequest
}

