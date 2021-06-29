package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分页拉取订单数据 API请求
alibaba.wdk.trade.order.balance.bill.query

提供接口供外部调用，分页拉取订单数据
*/
type AlibabaWdkTradeOrderBalanceBillQueryRequest struct {
    model.Params
    // 入参
    orderBalanceBillRequest   *OrderBalanceBillRequest
}

// 初始化AlibabaWdkTradeOrderBalanceBillQueryRequest对象
func NewAlibabaWdkTradeOrderBalanceBillQueryRequest() *AlibabaWdkTradeOrderBalanceBillQueryRequest{
    return &AlibabaWdkTradeOrderBalanceBillQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeOrderBalanceBillQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.order.balance.bill.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeOrderBalanceBillQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderBalanceBillRequest Setter
// 入参
func (r *AlibabaWdkTradeOrderBalanceBillQueryRequest) SetOrderBalanceBillRequest(orderBalanceBillRequest *OrderBalanceBillRequest) error {
    r.orderBalanceBillRequest = orderBalanceBillRequest
    r.Set("order_balance_bill_request", orderBalanceBillRequest)
    return nil
}

// OrderBalanceBillRequest Getter
func (r AlibabaWdkTradeOrderBalanceBillQueryRequest) GetOrderBalanceBillRequest() *OrderBalanceBillRequest {
    return r.orderBalanceBillRequest
}
