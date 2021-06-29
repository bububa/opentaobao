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
    _operation   string
    // 飞猪订单ID
    _orderId   string
    // 服务商ID
    _providerId   string
    // 司机服务状态。arriving-司机已出发，arrived-司机已到达，serviceStarted-已开始服务，serviceFinished-已结束服务
    _status   string
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
func (r *TaobaoAlitripCarOrderStatusRequest) SetOperation(_operation string) error {
    r._operation = _operation
    r.Set("operation", _operation)
    return nil
}

// Operation Getter
func (r TaobaoAlitripCarOrderStatusRequest) GetOperation() string {
    return r._operation
}
// OrderId Setter
// 飞猪订单ID
func (r *TaobaoAlitripCarOrderStatusRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoAlitripCarOrderStatusRequest) GetOrderId() string {
    return r._orderId
}
// ProviderId Setter
// 服务商ID
func (r *TaobaoAlitripCarOrderStatusRequest) SetProviderId(_providerId string) error {
    r._providerId = _providerId
    r.Set("provider_id", _providerId)
    return nil
}

// ProviderId Getter
func (r TaobaoAlitripCarOrderStatusRequest) GetProviderId() string {
    return r._providerId
}
// Status Setter
// 司机服务状态。arriving-司机已出发，arrived-司机已到达，serviceStarted-已开始服务，serviceFinished-已结束服务
func (r *TaobaoAlitripCarOrderStatusRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoAlitripCarOrderStatusRequest) GetStatus() string {
    return r._status
}
