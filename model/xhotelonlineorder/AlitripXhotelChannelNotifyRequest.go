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
type AlitripXhotelChannelNotifyAPIRequest struct {
    model.Params
    // 通知类型查询条件
    _orderNotifyQuery   *OrderNotifyQuery
}

// 初始化AlitripXhotelChannelNotifyAPIRequest对象
func NewAlitripXhotelChannelNotifyRequest() *AlitripXhotelChannelNotifyAPIRequest{
    return &AlitripXhotelChannelNotifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripXhotelChannelNotifyAPIRequest) GetApiMethodName() string {
    return "alitrip.xhotel.channel.notify"
}

// IRequest interface 方法, 获取API参数
func (r AlitripXhotelChannelNotifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderNotifyQuery Setter
// 通知类型查询条件
func (r *AlitripXhotelChannelNotifyAPIRequest) SetOrderNotifyQuery(_orderNotifyQuery *OrderNotifyQuery) error {
    r._orderNotifyQuery = _orderNotifyQuery
    r.Set("order_notify_query", _orderNotifyQuery)
    return nil
}

// OrderNotifyQuery Getter
func (r AlitripXhotelChannelNotifyAPIRequest) GetOrderNotifyQuery() *OrderNotifyQuery {
    return r._orderNotifyQuery
}
