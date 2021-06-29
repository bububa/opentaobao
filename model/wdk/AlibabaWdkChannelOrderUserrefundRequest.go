package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户发起售后退款(整单/部分) API请求
alibaba.wdk.channel.order.userrefund

用户发起售后退款(整单/部分)
*/
type AlibabaWdkChannelOrderUserrefundRequest struct {
    model.Params
    // 退款信息
    _orderUserRefundInfo   *OrderUserRefundInfo
}

// 初始化AlibabaWdkChannelOrderUserrefundRequest对象
func NewAlibabaWdkChannelOrderUserrefundRequest() *AlibabaWdkChannelOrderUserrefundRequest{
    return &AlibabaWdkChannelOrderUserrefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderUserrefundRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.order.userrefund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderUserrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderUserRefundInfo Setter
// 退款信息
func (r *AlibabaWdkChannelOrderUserrefundRequest) SetOrderUserRefundInfo(_orderUserRefundInfo *OrderUserRefundInfo) error {
    r._orderUserRefundInfo = _orderUserRefundInfo
    r.Set("order_user_refund_info", _orderUserRefundInfo)
    return nil
}

// OrderUserRefundInfo Getter
func (r AlibabaWdkChannelOrderUserrefundRequest) GetOrderUserRefundInfo() *OrderUserRefundInfo {
    return r._orderUserRefundInfo
}
