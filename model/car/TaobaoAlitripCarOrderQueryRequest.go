package car

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪订单状态查询接口 API请求
taobao.alitrip.car.order.query

提供给直连商家查询在飞猪平台上产生的订单
*/
type TaobaoAlitripCarOrderQueryRequest struct {
    model.Params
    // 飞猪平台订单id
    _orderId   string
}

// 初始化TaobaoAlitripCarOrderQueryRequest对象
func NewTaobaoAlitripCarOrderQueryRequest() *TaobaoAlitripCarOrderQueryRequest{
    return &TaobaoAlitripCarOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderQueryRequest) GetApiMethodName() string {
    return "taobao.alitrip.car.order.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 飞猪平台订单id
func (r *TaobaoAlitripCarOrderQueryRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoAlitripCarOrderQueryRequest) GetOrderId() string {
    return r._orderId
}
