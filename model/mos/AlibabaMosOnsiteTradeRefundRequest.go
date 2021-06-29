package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退款 API请求
alibaba.mos.onsite.trade.refund

当交易发生之后一段时间内，由于消费者或者商户的原因需退款，商户可通过退款接口将支付款退还给消费者，喵街将在收到退款请求并验证成功后，按退款规则将支付款按原路退到消费者账号上。

1. 交易超过可退款时间（签约时设置的可退款时间）的订单无法进行退款。
2. 只支持全额退款。
*/
type AlibabaMosOnsiteTradeRefundRequest struct {
    model.Params
    // 交易退款请求
    onsiteRefundRequest   *OnsiteRefundRequest
}

// 初始化AlibabaMosOnsiteTradeRefundRequest对象
func NewAlibabaMosOnsiteTradeRefundRequest() *AlibabaMosOnsiteTradeRefundRequest{
    return &AlibabaMosOnsiteTradeRefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosOnsiteTradeRefundRequest) GetApiMethodName() string {
    return "alibaba.mos.onsite.trade.refund"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosOnsiteTradeRefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OnsiteRefundRequest Setter
// 交易退款请求
func (r *AlibabaMosOnsiteTradeRefundRequest) SetOnsiteRefundRequest(onsiteRefundRequest *OnsiteRefundRequest) error {
    r.onsiteRefundRequest = onsiteRefundRequest
    r.Set("onsite_refund_request", onsiteRefundRequest)
    return nil
}

// OnsiteRefundRequest Getter
func (r AlibabaMosOnsiteTradeRefundRequest) GetOnsiteRefundRequest() *OnsiteRefundRequest {
    return r.onsiteRefundRequest
}
