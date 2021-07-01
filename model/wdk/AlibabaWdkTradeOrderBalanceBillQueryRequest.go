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
type AlibabaWdkTradeOrderBalanceBillQueryAPIRequest struct {
    model.Params
    // 入参
    _orderBalanceBillRequest   *OrderBalanceBillRequest
}

// 初始化AlibabaWdkTradeOrderBalanceBillQueryAPIRequest对象
func NewAlibabaWdkTradeOrderBalanceBillQueryRequest() *AlibabaWdkTradeOrderBalanceBillQueryAPIRequest{
    return &AlibabaWdkTradeOrderBalanceBillQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeOrderBalanceBillQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.trade.order.balance.bill.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeOrderBalanceBillQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderBalanceBillRequest Setter
// 入参
func (r *AlibabaWdkTradeOrderBalanceBillQueryAPIRequest) SetOrderBalanceBillRequest(_orderBalanceBillRequest *OrderBalanceBillRequest) error {
    r._orderBalanceBillRequest = _orderBalanceBillRequest
    r.Set("order_balance_bill_request", _orderBalanceBillRequest)
    return nil
}

// OrderBalanceBillRequest Getter
func (r AlibabaWdkTradeOrderBalanceBillQueryAPIRequest) GetOrderBalanceBillRequest() *OrderBalanceBillRequest {
    return r._orderBalanceBillRequest
}
