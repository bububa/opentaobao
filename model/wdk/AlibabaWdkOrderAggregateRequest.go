package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达订单按门店机台号聚合查询 APIRequest
alibaba.wdk.order.aggregate

淘鲜达订单按门店机台号聚合查询
*/
type AlibabaWdkOrderAggregateRequest struct {
    model.Params

    // 系统自动生成
    orderAggregateQueryRequest   *OrderAggregateQueryRequest 

}

func NewAlibabaWdkOrderAggregateRequest() *AlibabaWdkOrderAggregateRequest{
    return &AlibabaWdkOrderAggregateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkOrderAggregateRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.aggregate"
}

func (r AlibabaWdkOrderAggregateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkOrderAggregateRequest) SetOrderAggregateQueryRequest(orderAggregateQueryRequest *OrderAggregateQueryRequest) error {
    r.orderAggregateQueryRequest = orderAggregateQueryRequest
    r.Set("order_aggregate_query_request", orderAggregateQueryRequest)
    return nil
}

func (r AlibabaWdkOrderAggregateRequest) GetOrderAggregateQueryRequest() *OrderAggregateQueryRequest {
    return r.orderAggregateQueryRequest
}

