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
type AlibabaWdkChannelOrderCreateRequest struct {
    model.Params
    // 订单信息
    orderInfo   *OrderInfo
}

// 初始化AlibabaWdkChannelOrderCreateRequest对象
func NewAlibabaWdkChannelOrderCreateRequest() *AlibabaWdkChannelOrderCreateRequest{
    return &AlibabaWdkChannelOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderInfo Setter
// 订单信息
func (r *AlibabaWdkChannelOrderCreateRequest) SetOrderInfo(orderInfo *OrderInfo) error {
    r.orderInfo = orderInfo
    r.Set("order_info", orderInfo)
    return nil
}

// OrderInfo Getter
func (r AlibabaWdkChannelOrderCreateRequest) GetOrderInfo() *OrderInfo {
    return r.orderInfo
}
