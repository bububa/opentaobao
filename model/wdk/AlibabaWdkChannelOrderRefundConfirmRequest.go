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
type AlibabaWdkChannelOrderRefundConfirmAPIRequest struct {
    model.Params
    // 退款确认信息
    _orderRefundConfirmInfo   *OrderRefundConfirmInfo
}

// 初始化AlibabaWdkChannelOrderRefundConfirmAPIRequest对象
func NewAlibabaWdkChannelOrderRefundConfirmRequest() *AlibabaWdkChannelOrderRefundConfirmAPIRequest{
    return &AlibabaWdkChannelOrderRefundConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderRefundConfirmAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.order.refund.confirm"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderRefundConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderRefundConfirmInfo Setter
// 退款确认信息
func (r *AlibabaWdkChannelOrderRefundConfirmAPIRequest) SetOrderRefundConfirmInfo(_orderRefundConfirmInfo *OrderRefundConfirmInfo) error {
    r._orderRefundConfirmInfo = _orderRefundConfirmInfo
    r.Set("order_refund_confirm_info", _orderRefundConfirmInfo)
    return nil
}

// OrderRefundConfirmInfo Getter
func (r AlibabaWdkChannelOrderRefundConfirmAPIRequest) GetOrderRefundConfirmInfo() *OrderRefundConfirmInfo {
    return r._orderRefundConfirmInfo
}
