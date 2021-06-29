package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 订单创建 API请求
alibaba.tcls.aelophy.merchant.channel.order.create

翱象小程序渠道订单创建
*/
type AlibabaTclsAelophyMerchantChannelOrderCreateRequest struct {
    model.Params
    // 订单信息
    _orderInfo   *OrderInfo
}

// 初始化AlibabaTclsAelophyMerchantChannelOrderCreateRequest对象
func NewAlibabaTclsAelophyMerchantChannelOrderCreateRequest() *AlibabaTclsAelophyMerchantChannelOrderCreateRequest{
    return &AlibabaTclsAelophyMerchantChannelOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.channel.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderInfo Setter
// 订单信息
func (r *AlibabaTclsAelophyMerchantChannelOrderCreateRequest) SetOrderInfo(_orderInfo *OrderInfo) error {
    r._orderInfo = _orderInfo
    r.Set("order_info", _orderInfo)
    return nil
}

// OrderInfo Getter
func (r AlibabaTclsAelophyMerchantChannelOrderCreateRequest) GetOrderInfo() *OrderInfo {
    return r._orderInfo
}
