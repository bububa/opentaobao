package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOnsiteTradeRefundAPIRequest 退款 API请求
// alibaba.mos.onsite.trade.refund
//
// 当交易发生之后一段时间内，由于消费者或者商户的原因需退款，商户可通过退款接口将支付款退还给消费者，喵街将在收到退款请求并验证成功后，按退款规则将支付款按原路退到消费者账号上。
//
// 1. 交易超过可退款时间（签约时设置的可退款时间）的订单无法进行退款。
// 2. 只支持全额退款。
type AlibabaMosOnsiteTradeRefundAPIRequest struct {
	model.Params
	// 交易退款请求
	_onsiteRefundRequest *OnsiteRefundRequest
}

// NewAlibabaMosOnsiteTradeRefundRequest 初始化AlibabaMosOnsiteTradeRefundAPIRequest对象
func NewAlibabaMosOnsiteTradeRefundRequest() *AlibabaMosOnsiteTradeRefundAPIRequest {
	return &AlibabaMosOnsiteTradeRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosOnsiteTradeRefundAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.onsite.trade.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosOnsiteTradeRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMosOnsiteTradeRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOnsiteRefundRequest is OnsiteRefundRequest Setter
// 交易退款请求
func (r *AlibabaMosOnsiteTradeRefundAPIRequest) SetOnsiteRefundRequest(_onsiteRefundRequest *OnsiteRefundRequest) error {
	r._onsiteRefundRequest = _onsiteRefundRequest
	r.Set("onsite_refund_request", _onsiteRefundRequest)
	return nil
}

// GetOnsiteRefundRequest OnsiteRefundRequest Getter
func (r AlibabaMosOnsiteTradeRefundAPIRequest) GetOnsiteRefundRequest() *OnsiteRefundRequest {
	return r._onsiteRefundRequest
}
