package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOnsiteTradeIsnewpayorderAPIRequest 是否为新支付订单 API请求
// alibaba.mos.onsite.trade.isnewpayorder
//
// 退款时，老支付宝手淘退款接口需要查一下该订单是否为新支付订单
type AlibabaMosOnsiteTradeIsnewpayorderAPIRequest struct {
	model.Params
	// 外部订单号
	_outTradeNo string
}

// NewAlibabaMosOnsiteTradeIsnewpayorderRequest 初始化AlibabaMosOnsiteTradeIsnewpayorderAPIRequest对象
func NewAlibabaMosOnsiteTradeIsnewpayorderRequest() *AlibabaMosOnsiteTradeIsnewpayorderAPIRequest {
	return &AlibabaMosOnsiteTradeIsnewpayorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosOnsiteTradeIsnewpayorderAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.onsite.trade.isnewpayorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosOnsiteTradeIsnewpayorderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOutTradeNo is OutTradeNo Setter
// 外部订单号
func (r *AlibabaMosOnsiteTradeIsnewpayorderAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// GetOutTradeNo OutTradeNo Getter
func (r AlibabaMosOnsiteTradeIsnewpayorderAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}
