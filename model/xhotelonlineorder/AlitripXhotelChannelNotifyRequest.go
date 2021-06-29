package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
分销渠道各类通知接口 API请求
alitrip.xhotel.channel.notify

分销渠道支付通知
*/
type AlitripXhotelChannelNotifyRequest struct {
    model.Params
    // 通知类型查询条件
    _orderNotifyQuery   *OrderNotifyQuery
}

// 初始化AlitripXhotelChannelNotifyRequest对象
func NewAlitripXhotelChannelNotifyRequest() *AlitripXhotelChannelNotifyRequest{
    return &AlitripXhotelChannelNotifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripXhotelChannelNotifyRequest) GetApiMethodName() string {
    return "alitrip.xhotel.channel.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlitripXhotelChannelNotifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderNotifyQuery Setter
// 通知类型查询条件
func (r *AlitripXhotelChannelNotifyRequest) SetOrderNotifyQuery(_orderNotifyQuery *OrderNotifyQuery) error {
    r._orderNotifyQuery = _orderNotifyQuery
    r.Set("order_notify_query", _orderNotifyQuery)
    return nil
}

// OrderNotifyQuery Getter
func (r AlitripXhotelChannelNotifyRequest) GetOrderNotifyQuery() *OrderNotifyQuery {
    return r._orderNotifyQuery
}
