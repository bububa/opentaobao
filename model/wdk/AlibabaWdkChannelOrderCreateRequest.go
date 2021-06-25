package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 APIRequest
alibaba.wdk.channel.order.create

外部商家创建订单
*/
type AlibabaWdkChannelOrderCreateRequest struct {
    model.Params

    // 订单信息
    orderInfo   *OrderInfo 

}

func NewAlibabaWdkChannelOrderCreateRequest() *AlibabaWdkChannelOrderCreateRequest{
    return &AlibabaWdkChannelOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkChannelOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.order.create"
}

func (r AlibabaWdkChannelOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkChannelOrderCreateRequest) SetOrderInfo(orderInfo *OrderInfo) error {
    r.orderInfo = orderInfo
    r.Set("order_info", orderInfo)
    return nil
}

func (r AlibabaWdkChannelOrderCreateRequest) GetOrderInfo() *OrderInfo {
    return r.orderInfo
}

