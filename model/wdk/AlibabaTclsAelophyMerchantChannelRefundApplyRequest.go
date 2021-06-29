package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
翱象商家自有渠道 逆向单申请 API请求
alibaba.tcls.aelophy.merchant.channel.refund.apply

翱象小程序 用户逆向单申请
*/
type AlibabaTclsAelophyMerchantChannelRefundApplyRequest struct {
    model.Params
    // 请求对象
    _refundApplyInfo   *RefundApplyInfo
}

// 初始化AlibabaTclsAelophyMerchantChannelRefundApplyRequest对象
func NewAlibabaTclsAelophyMerchantChannelRefundApplyRequest() *AlibabaTclsAelophyMerchantChannelRefundApplyRequest{
    return &AlibabaTclsAelophyMerchantChannelRefundApplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelRefundApplyRequest) GetApiMethodName() string {
    return "alibaba.tcls.aelophy.merchant.channel.refund.apply"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelRefundApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundApplyInfo Setter
// 请求对象
func (r *AlibabaTclsAelophyMerchantChannelRefundApplyRequest) SetRefundApplyInfo(_refundApplyInfo *RefundApplyInfo) error {
    r._refundApplyInfo = _refundApplyInfo
    r.Set("refund_apply_info", _refundApplyInfo)
    return nil
}

// RefundApplyInfo Getter
func (r AlibabaTclsAelophyMerchantChannelRefundApplyRequest) GetRefundApplyInfo() *RefundApplyInfo {
    return r._refundApplyInfo
}
