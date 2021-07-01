package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认签收 API请求
alibaba.idle.rent.order.receiveitem

确认揽收/签收
*/
type AlibabaIdleRentOrderReceiveitemAPIRequest struct {
    model.Params
    // 订单id
    _orderId   int64
}

// 初始化AlibabaIdleRentOrderReceiveitemAPIRequest对象
func NewAlibabaIdleRentOrderReceiveitemRequest() *AlibabaIdleRentOrderReceiveitemAPIRequest{
    return &AlibabaIdleRentOrderReceiveitemAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderReceiveitemAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.receiveitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderReceiveitemAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderReceiveitemAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleRentOrderReceiveitemAPIRequest) GetOrderId() int64 {
    return r._orderId
}
