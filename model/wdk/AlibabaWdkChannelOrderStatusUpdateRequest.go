package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单状态变更 API请求
alibaba.wdk.channel.order.status.update

订单状态变更
*/
type AlibabaWdkChannelOrderStatusUpdateAPIRequest struct {
    model.Params
    // 修改信息
    _orderStatusInfo   *OrderStatusInfo
}

// 初始化AlibabaWdkChannelOrderStatusUpdateAPIRequest对象
func NewAlibabaWdkChannelOrderStatusUpdateRequest() *AlibabaWdkChannelOrderStatusUpdateAPIRequest{
    return &AlibabaWdkChannelOrderStatusUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderStatusUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.order.status.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderStatusUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderStatusInfo Setter
// 修改信息
func (r *AlibabaWdkChannelOrderStatusUpdateAPIRequest) SetOrderStatusInfo(_orderStatusInfo *OrderStatusInfo) error {
    r._orderStatusInfo = _orderStatusInfo
    r.Set("order_status_info", _orderStatusInfo)
    return nil
}

// OrderStatusInfo Getter
func (r AlibabaWdkChannelOrderStatusUpdateAPIRequest) GetOrderStatusInfo() *OrderStatusInfo {
    return r._orderStatusInfo
}
