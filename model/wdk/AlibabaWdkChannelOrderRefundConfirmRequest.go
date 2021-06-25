package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款确认 APIRequest
alibaba.wdk.channel.order.refund.confirm

退款确认
*/
type AlibabaWdkChannelOrderRefundConfirmRequest struct {
    model.Params

    // 退款确认信息
    orderRefundConfirmInfo   *OrderRefundConfirmInfo 

}

func NewAlibabaWdkChannelOrderRefundConfirmRequest() *AlibabaWdkChannelOrderRefundConfirmRequest{
    return &AlibabaWdkChannelOrderRefundConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkChannelOrderRefundConfirmRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.order.refund.confirm"
}

func (r AlibabaWdkChannelOrderRefundConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkChannelOrderRefundConfirmRequest) SetOrderRefundConfirmInfo(orderRefundConfirmInfo *OrderRefundConfirmInfo) error {
    r.orderRefundConfirmInfo = orderRefundConfirmInfo
    r.Set("order_refund_confirm_info", orderRefundConfirmInfo)
    return nil
}

func (r AlibabaWdkChannelOrderRefundConfirmRequest) GetOrderRefundConfirmInfo() *OrderRefundConfirmInfo {
    return r.orderRefundConfirmInfo
}

