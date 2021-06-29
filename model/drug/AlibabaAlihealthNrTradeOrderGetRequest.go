package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单详情 API请求
alibaba.alihealth.nr.trade.order.get

阿里健康O2O，获取订单详情
*/
type AlibabaAlihealthNrTradeOrderGetRequest struct {
    model.Params
    // 淘宝订单ID
    _orderId   int64
}

// 初始化AlibabaAlihealthNrTradeOrderGetRequest对象
func NewAlibabaAlihealthNrTradeOrderGetRequest() *AlibabaAlihealthNrTradeOrderGetRequest{
    return &AlibabaAlihealthNrTradeOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrTradeOrderGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nr.trade.order.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrTradeOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 淘宝订单ID
func (r *AlibabaAlihealthNrTradeOrderGetRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthNrTradeOrderGetRequest) GetOrderId() int64 {
    return r._orderId
}
