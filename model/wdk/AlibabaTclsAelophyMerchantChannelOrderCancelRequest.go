package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 交易订单取消 APIRequest
alibaba.tcls.aelophy.merchant.channel.order.cancel

翱象小程序用户取消订单
*/
type AlibabaTclsAelophyMerchantChannelOrderCancelRequest struct {
    model.Params

    // 取消信息
    userCancelInfo   *OrderUserCancelInfo 

}

func NewAlibabaTclsAelophyMerchantChannelOrderCancelRequest() *AlibabaTclsAelophyMerchantChannelOrderCancelRequest{
    return &AlibabaTclsAelophyMerchantChannelOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsAelophyMerchantChannelOrderCancelRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.channel.order.cancel"
}

func (r AlibabaTclsAelophyMerchantChannelOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsAelophyMerchantChannelOrderCancelRequest) SetUserCancelInfo(userCancelInfo *OrderUserCancelInfo) error {
    r.userCancelInfo = userCancelInfo
    r.Set("user_cancel_info", userCancelInfo)
    return nil
}

func (r AlibabaTclsAelophyMerchantChannelOrderCancelRequest) GetUserCancelInfo() *OrderUserCancelInfo {
    return r.userCancelInfo
}

