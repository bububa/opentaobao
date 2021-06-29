package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 交易订单取消 API请求
alibaba.tcls.aelophy.merchant.channel.order.cancel

翱象小程序用户取消订单
*/
type AlibabaTclsAelophyMerchantChannelOrderCancelRequest struct {
    model.Params
    // 取消信息
    userCancelInfo   *OrderUserCancelInfo
}

// 初始化AlibabaTclsAelophyMerchantChannelOrderCancelRequest对象
func NewAlibabaTclsAelophyMerchantChannelOrderCancelRequest() *AlibabaTclsAelophyMerchantChannelOrderCancelRequest{
    return &AlibabaTclsAelophyMerchantChannelOrderCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelOrderCancelRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.channel.order.cancel"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserCancelInfo Setter
// 取消信息
func (r *AlibabaTclsAelophyMerchantChannelOrderCancelRequest) SetUserCancelInfo(userCancelInfo *OrderUserCancelInfo) error {
    r.userCancelInfo = userCancelInfo
    r.Set("user_cancel_info", userCancelInfo)
    return nil
}

// UserCancelInfo Getter
func (r AlibabaTclsAelophyMerchantChannelOrderCancelRequest) GetUserCancelInfo() *OrderUserCancelInfo {
    return r.userCancelInfo
}
