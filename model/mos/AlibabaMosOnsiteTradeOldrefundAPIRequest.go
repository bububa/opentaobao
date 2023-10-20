package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosonsitetradeoldrefundAPIRequest 线下新退款接口（专为老退款接口调用） API请求
// alibaba.mos.onsite.trade.oldrefund
//
// 线下新退款接口（专为老退款接口调用）。新接口支付支付宝、手淘、天猫，老接口退款时，需要调用该接口退新单，并适配老接口响应参数返回
type AlibabamosonsitetradeoldrefundAPIRequest struct {
	model.Params
	// 交易退款请求
	_onsiteRefundRequest *OnsiteRefundRequest
}

// NewAlibabamosonsitetradeoldrefundRequest 初始化AlibabamosonsitetradeoldrefundAPIRequest对象
func NewAlibabamosonsitetradeoldrefundRequest() *AlibabamosonsitetradeoldrefundAPIRequest {
	return &AlibabamosonsitetradeoldrefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosonsitetradeoldrefundAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.onsite.trade.oldrefund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosonsitetradeoldrefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosonsitetradeoldrefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOnsiteRefundRequest is OnsiteRefundRequest Setter
// 交易退款请求
func (r *AlibabamosonsitetradeoldrefundAPIRequest) SetOnsiteRefundRequest(_onsiteRefundRequest *OnsiteRefundRequest) error {
	r._onsiteRefundRequest = _onsiteRefundRequest
	r.Set("onsite_refund_request", _onsiteRefundRequest)
	return nil
}

// GetOnsiteRefundRequest OnsiteRefundRequest Getter
func (r AlibabamosonsitetradeoldrefundAPIRequest) GetOnsiteRefundRequest() *OnsiteRefundRequest {
	return r._onsiteRefundRequest
}
