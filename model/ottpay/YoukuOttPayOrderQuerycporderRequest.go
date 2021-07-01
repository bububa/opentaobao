package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询支付订单对应cp订单号 API请求
youku.ott.pay.order.querycporder

根据支付订单查询对应cp订单号
*/
type YoukuOttPayOrderQuerycporderAPIRequest struct {
    model.Params
    // 支付对应订单
    _gatewayOrder   string
}

// 初始化YoukuOttPayOrderQuerycporderAPIRequest对象
func NewYoukuOttPayOrderQuerycporderRequest() *YoukuOttPayOrderQuerycporderAPIRequest{
    return &YoukuOttPayOrderQuerycporderAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderQuerycporderAPIRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.querycporder"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderQuerycporderAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GatewayOrder Setter
// 支付对应订单
func (r *YoukuOttPayOrderQuerycporderAPIRequest) SetGatewayOrder(_gatewayOrder string) error {
    r._gatewayOrder = _gatewayOrder
    r.Set("gateway_order", _gatewayOrder)
    return nil
}

// GatewayOrder Getter
func (r YoukuOttPayOrderQuerycporderAPIRequest) GetGatewayOrder() string {
    return r._gatewayOrder
}
