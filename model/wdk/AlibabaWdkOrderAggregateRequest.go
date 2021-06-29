package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘鲜达订单按门店机台号聚合查询 API请求
alibaba.wdk.order.aggregate

淘鲜达订单按门店机台号聚合查询
*/
type AlibabaWdkOrderAggregateRequest struct {
    model.Params
    // 系统自动生成
    orderAggregateQueryRequest   *OrderAggregateQueryRequest
}

// 初始化AlibabaWdkOrderAggregateRequest对象
func NewAlibabaWdkOrderAggregateRequest() *AlibabaWdkOrderAggregateRequest{
    return &AlibabaWdkOrderAggregateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderAggregateRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.aggregate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderAggregateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderAggregateQueryRequest Setter
// 系统自动生成
func (r *AlibabaWdkOrderAggregateRequest) SetOrderAggregateQueryRequest(orderAggregateQueryRequest *OrderAggregateQueryRequest) error {
    r.orderAggregateQueryRequest = orderAggregateQueryRequest
    r.Set("order_aggregate_query_request", orderAggregateQueryRequest)
    return nil
}

// OrderAggregateQueryRequest Getter
func (r AlibabaWdkOrderAggregateRequest) GetOrderAggregateQueryRequest() *OrderAggregateQueryRequest {
    return r.orderAggregateQueryRequest
}
