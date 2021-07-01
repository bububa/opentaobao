package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下新退款接口（专为老退款接口调用） API请求
alibaba.mos.onsite.trade.oldrefund

线下新退款接口（专为老退款接口调用）。新接口支付支付宝、手淘、天猫，老接口退款时，需要调用该接口退新单，并适配老接口响应参数返回
*/
type AlibabaMosOnsiteTradeOldrefundAPIRequest struct {
    model.Params
    // 交易退款请求
    _onsiteRefundRequest   *OnsiteRefundRequest
}

// 初始化AlibabaMosOnsiteTradeOldrefundAPIRequest对象
func NewAlibabaMosOnsiteTradeOldrefundRequest() *AlibabaMosOnsiteTradeOldrefundAPIRequest{
    return &AlibabaMosOnsiteTradeOldrefundAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosOnsiteTradeOldrefundAPIRequest) GetApiMethodName() string {
    return "alibaba.mos.onsite.trade.oldrefund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosOnsiteTradeOldrefundAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OnsiteRefundRequest Setter
// 交易退款请求
func (r *AlibabaMosOnsiteTradeOldrefundAPIRequest) SetOnsiteRefundRequest(_onsiteRefundRequest *OnsiteRefundRequest) error {
    r._onsiteRefundRequest = _onsiteRefundRequest
    r.Set("onsite_refund_request", _onsiteRefundRequest)
    return nil
}

// OnsiteRefundRequest Getter
func (r AlibabaMosOnsiteTradeOldrefundAPIRequest) GetOnsiteRefundRequest() *OnsiteRefundRequest {
    return r._onsiteRefundRequest
}
