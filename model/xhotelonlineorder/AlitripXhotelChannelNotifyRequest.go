package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销渠道各类通知接口 APIRequest
alitrip.xhotel.channel.notify

分销渠道支付通知
*/
type AlitripXhotelChannelNotifyRequest struct {
    model.Params

    // 通知类型查询条件
    orderNotifyQuery   *OrderNotifyQuery 

}

func NewAlitripXhotelChannelNotifyRequest() *AlitripXhotelChannelNotifyRequest{
    return &AlitripXhotelChannelNotifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripXhotelChannelNotifyRequest) GetApiMethodName() string {
    return "alitrip.xhotel.channel.notify"
}

func (r AlitripXhotelChannelNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripXhotelChannelNotifyRequest) SetOrderNotifyQuery(orderNotifyQuery *OrderNotifyQuery) error {
    r.orderNotifyQuery = orderNotifyQuery
    r.Set("order_notify_query", orderNotifyQuery)
    return nil
}

func (r AlitripXhotelChannelNotifyRequest) GetOrderNotifyQuery() *OrderNotifyQuery {
    return r.orderNotifyQuery
}

