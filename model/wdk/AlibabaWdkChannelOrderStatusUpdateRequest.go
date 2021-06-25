package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单状态变更 APIRequest
alibaba.wdk.channel.order.status.update

订单状态变更
*/
type AlibabaWdkChannelOrderStatusUpdateRequest struct {
    model.Params

    // 修改信息
    orderStatusInfo   *OrderStatusInfo 

}

func NewAlibabaWdkChannelOrderStatusUpdateRequest() *AlibabaWdkChannelOrderStatusUpdateRequest{
    return &AlibabaWdkChannelOrderStatusUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkChannelOrderStatusUpdateRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.order.status.update"
}

func (r AlibabaWdkChannelOrderStatusUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkChannelOrderStatusUpdateRequest) SetOrderStatusInfo(orderStatusInfo *OrderStatusInfo) error {
    r.orderStatusInfo = orderStatusInfo
    r.Set("order_status_info", orderStatusInfo)
    return nil
}

func (r AlibabaWdkChannelOrderStatusUpdateRequest) GetOrderStatusInfo() *OrderStatusInfo {
    return r.orderStatusInfo
}

