package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家订单状态改变通知接口（神州专车接口） API请求
taobao.alitrip.car.order.status

商家订单状态改变通知接口，神州专车专用接口！
*/
type TaobaoAlitripCarOrderStatusRequest struct {
    model.Params
    // 固定值：statusChanged
    operation   string
    // 飞猪订单ID
    orderId   string
    // 服务商ID
    providerId   string
    // 司机服务状态。arriving-司机已出发，arrived-司机已到达，serviceStarted-已开始服务，serviceFinished-已结束服务
    status   string
}

// 初始化TaobaoAlitripCarOrderStatusRequest对象
func NewTaobaoAlitripCarOrderStatusRequest() *TaobaoAlitripCarOrderStatusRequest{
    return &TaobaoAlitripCarOrderStatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderStatusRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.order.status"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderStatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Operation Setter
// 固定值：statusChanged
func (r *TaobaoAlitripCarOrderStatusRequest) SetOperation(operation string) error {
    r.operation = operation
    r.Set("operation", operation)
    return nil
}

// Operation Getter
func (r TaobaoAlitripCarOrderStatusRequest) GetOperation() string {
    return r.operation
}
// OrderId Setter
// 飞猪订单ID
func (r *TaobaoAlitripCarOrderStatusRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoAlitripCarOrderStatusRequest) GetOrderId() string {
    return r.orderId
}
// ProviderId Setter
// 服务商ID
func (r *TaobaoAlitripCarOrderStatusRequest) SetProviderId(providerId string) error {
    r.providerId = providerId
    r.Set("provider_id", providerId)
    return nil
}

// ProviderId Getter
func (r TaobaoAlitripCarOrderStatusRequest) GetProviderId() string {
    return r.providerId
}
// Status Setter
// 司机服务状态。arriving-司机已出发，arrived-司机已到达，serviceStarted-已开始服务，serviceFinished-已结束服务
func (r *TaobaoAlitripCarOrderStatusRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoAlitripCarOrderStatusRequest) GetStatus() string {
    return r.status
}
