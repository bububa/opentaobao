package ottpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询支付订单对应cp订单号 APIRequest
youku.ott.pay.order.querycporder

根据支付订单查询对应cp订单号
*/
type YoukuOttPayOrderQuerycporderRequest struct {
    model.Params

    // 支付对应订单
    gatewayOrder   string 

}

func NewYoukuOttPayOrderQuerycporderRequest() *YoukuOttPayOrderQuerycporderRequest{
    return &YoukuOttPayOrderQuerycporderRequest{
        Params: model.NewParams(),
    }
}

func (r YoukuOttPayOrderQuerycporderRequest) GetApiMethodName() string {
    return "youku.ott.pay.order.querycporder"
}

func (r YoukuOttPayOrderQuerycporderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YoukuOttPayOrderQuerycporderRequest) SetGatewayOrder(gatewayOrder string) error {
    r.gatewayOrder = gatewayOrder
    r.Set("gateway_order", gatewayOrder)
    return nil
}

func (r YoukuOttPayOrderQuerycporderRequest) GetGatewayOrder() string {
    return r.gatewayOrder
}

