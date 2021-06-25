package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 订单状态更新 APIRequest
alibaba.tcls.aelophy.merchant.channel.order.updatestatus

订单状态变更
*/
type AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest struct {
    model.Params

    // 修改信息
    orderStatusInfo   *OrderStatusInfo 

}

func NewAlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest() *AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest{
    return &AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.channel.order.updatestatus"
}

func (r AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest) SetOrderStatusInfo(orderStatusInfo *OrderStatusInfo) error {
    r.orderStatusInfo = orderStatusInfo
    r.Set("order_status_info", orderStatusInfo)
    return nil
}

func (r AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest) GetOrderStatusInfo() *OrderStatusInfo {
    return r.orderStatusInfo
}

