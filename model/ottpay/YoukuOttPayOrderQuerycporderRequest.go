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
type YoukuOttPayOrderQuerycporderRequest struct {
    model.Params
    // 支付对应订单
    _gatewayOrder   string
}

// 初始化YoukuOttPayOrderQuerycporderRequest对象
func NewYoukuOttPayOrderQuerycporderRequest() *YoukuOttPayOrderQuerycporderRequest{
    return &YoukuOttPayOrderQuerycporderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderQuerycporderRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.querycporder"
}

// IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderQuerycporderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GatewayOrder Setter
// 支付对应订单
func (r *YoukuOttPayOrderQuerycporderRequest) SetGatewayOrder(_gatewayOrder string) error {
    r._gatewayOrder = _gatewayOrder
    r.Set("gateway_order", _gatewayOrder)
    return nil
}

// GatewayOrder Getter
func (r YoukuOttPayOrderQuerycporderRequest) GetGatewayOrder() string {
    return r._gatewayOrder
}
