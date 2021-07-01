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
type AlibabaWdkOrderAggregateAPIRequest struct {
    model.Params
    // 系统自动生成
    _orderAggregateQueryRequest   *OrderAggregateQueryRequest
}

// 初始化AlibabaWdkOrderAggregateAPIRequest对象
func NewAlibabaWdkOrderAggregateRequest() *AlibabaWdkOrderAggregateAPIRequest{
    return &AlibabaWdkOrderAggregateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkOrderAggregateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.aggregate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkOrderAggregateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderAggregateQueryRequest Setter
// 系统自动生成
func (r *AlibabaWdkOrderAggregateAPIRequest) SetOrderAggregateQueryRequest(_orderAggregateQueryRequest *OrderAggregateQueryRequest) error {
    r._orderAggregateQueryRequest = _orderAggregateQueryRequest
    r.Set("order_aggregate_query_request", _orderAggregateQueryRequest)
    return nil
}

// OrderAggregateQueryRequest Getter
func (r AlibabaWdkOrderAggregateAPIRequest) GetOrderAggregateQueryRequest() *OrderAggregateQueryRequest {
    return r._orderAggregateQueryRequest
}
