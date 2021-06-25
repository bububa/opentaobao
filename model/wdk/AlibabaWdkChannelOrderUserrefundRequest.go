package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户发起售后退款(整单/部分) APIRequest
alibaba.wdk.channel.order.userrefund

用户发起售后退款(整单/部分)
*/
type AlibabaWdkChannelOrderUserrefundRequest struct {
    model.Params

    // 退款信息
    orderUserRefundInfo   *OrderUserRefundInfo 

}

func NewAlibabaWdkChannelOrderUserrefundRequest() *AlibabaWdkChannelOrderUserrefundRequest{
    return &AlibabaWdkChannelOrderUserrefundRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkChannelOrderUserrefundRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.order.userrefund"
}

func (r AlibabaWdkChannelOrderUserrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkChannelOrderUserrefundRequest) SetOrderUserRefundInfo(orderUserRefundInfo *OrderUserRefundInfo) error {
    r.orderUserRefundInfo = orderUserRefundInfo
    r.Set("order_user_refund_info", orderUserRefundInfo)
    return nil
}

func (r AlibabaWdkChannelOrderUserrefundRequest) GetOrderUserRefundInfo() *OrderUserRefundInfo {
    return r.orderUserRefundInfo
}

