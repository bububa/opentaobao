package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 API请求
alibaba.wdk.channel.order.create

外部商家创建订单
*/
type AlibabaWdkChannelOrderCreateAPIRequest struct {
    model.Params
    // 订单信息
    _orderInfo   *OrderInfo
}

// 初始化AlibabaWdkChannelOrderCreateAPIRequest对象
func NewAlibabaWdkChannelOrderCreateRequest() *AlibabaWdkChannelOrderCreateAPIRequest{
    return &AlibabaWdkChannelOrderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderInfo Setter
// 订单信息
func (r *AlibabaWdkChannelOrderCreateAPIRequest) SetOrderInfo(_orderInfo *OrderInfo) error {
    r._orderInfo = _orderInfo
    r.Set("order_info", _orderInfo)
    return nil
}

// OrderInfo Getter
func (r AlibabaWdkChannelOrderCreateAPIRequest) GetOrderInfo() *OrderInfo {
    return r._orderInfo
}
