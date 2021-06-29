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
type AlibabaIdleRentOrderReceiveitemRequest struct {
    model.Params
    // 订单id
    orderId   int64
}

// 初始化AlibabaIdleRentOrderReceiveitemRequest对象
func NewAlibabaIdleRentOrderReceiveitemRequest() *AlibabaIdleRentOrderReceiveitemRequest{
    return &AlibabaIdleRentOrderReceiveitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderReceiveitemRequest) GetApiMethodName() string {
    return "alibaba.idle.rent.order.receiveitem"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderReceiveitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderReceiveitemRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaIdleRentOrderReceiveitemRequest) GetOrderId() int64 {
    return r.orderId
}
