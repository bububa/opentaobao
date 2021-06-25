package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 订单创建 APIRequest
alibaba.tcls.aelophy.merchant.channel.order.create

翱象小程序渠道订单创建
*/
type AlibabaTclsAelophyMerchantChannelOrderCreateRequest struct {
    model.Params

    // 订单信息
    orderInfo   *OrderInfo 

}

func NewAlibabaTclsAelophyMerchantChannelOrderCreateRequest() *AlibabaTclsAelophyMerchantChannelOrderCreateRequest{
    return &AlibabaTclsAelophyMerchantChannelOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyMerchantChannelOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.channel.order.create"
}

func (r AlibabaTclsAelophyMerchantChannelOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyMerchantChannelOrderCreateRequest) SetOrderInfo(orderInfo *OrderInfo) error {
    r.orderInfo = orderInfo
    r.Set("order_info", orderInfo)
    return nil
}

func (r AlibabaTclsAelophyMerchantChannelOrderCreateRequest) GetOrderInfo() *OrderInfo {
    return r.orderInfo
}

