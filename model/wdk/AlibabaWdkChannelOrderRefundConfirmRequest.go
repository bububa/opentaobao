package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款确认 API请求
alibaba.wdk.channel.order.refund.confirm

退款确认
*/
type AlibabaWdkChannelOrderRefundConfirmRequest struct {
    model.Params
    // 退款确认信息
    orderRefundConfirmInfo   *OrderRefundConfirmInfo
}

// 初始化AlibabaWdkChannelOrderRefundConfirmRequest对象
func NewAlibabaWdkChannelOrderRefundConfirmRequest() *AlibabaWdkChannelOrderRefundConfirmRequest{
    return &AlibabaWdkChannelOrderRefundConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderRefundConfirmRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.order.refund.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderRefundConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderRefundConfirmInfo Setter
// 退款确认信息
func (r *AlibabaWdkChannelOrderRefundConfirmRequest) SetOrderRefundConfirmInfo(orderRefundConfirmInfo *OrderRefundConfirmInfo) error {
    r.orderRefundConfirmInfo = orderRefundConfirmInfo
    r.Set("order_refund_confirm_info", orderRefundConfirmInfo)
    return nil
}

// OrderRefundConfirmInfo Getter
func (r AlibabaWdkChannelOrderRefundConfirmRequest) GetOrderRefundConfirmInfo() *OrderRefundConfirmInfo {
    return r.orderRefundConfirmInfo
}
