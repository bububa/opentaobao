package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 订单状态更新 API请求
alibaba.tcls.aelophy.merchant.channel.order.updatestatus

订单状态变更
*/
type AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest struct {
    model.Params
    // 修改信息
    orderStatusInfo   *OrderStatusInfo
}

// 初始化AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest对象
func NewAlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest() *AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest{
    return &AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.channel.order.updatestatus"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderStatusInfo Setter
// 修改信息
func (r *AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest) SetOrderStatusInfo(orderStatusInfo *OrderStatusInfo) error {
    r.orderStatusInfo = orderStatusInfo
    r.Set("order_status_info", orderStatusInfo)
    return nil
}

// OrderStatusInfo Getter
func (r AlibabaTclsAelophyMerchantChannelOrderUpdatestatusRequest) GetOrderStatusInfo() *OrderStatusInfo {
    return r.orderStatusInfo
}
