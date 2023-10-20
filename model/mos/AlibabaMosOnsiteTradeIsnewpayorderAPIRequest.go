package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosonsitetradeisnewpayorderAPIRequest 是否为新支付订单 API请求
// alibaba.mos.onsite.trade.isnewpayorder
//
// 退款时，老支付宝手淘退款接口需要查一下该订单是否为新支付订单
type AlibabamosonsitetradeisnewpayorderAPIRequest struct {
	model.Params
	// 外部订单号
	_outTradeNo string
}

// NewAlibabamosonsitetradeisnewpayorderRequest 初始化AlibabamosonsitetradeisnewpayorderAPIRequest对象
func NewAlibabamosonsitetradeisnewpayorderRequest() *AlibabamosonsitetradeisnewpayorderAPIRequest {
	return &AlibabamosonsitetradeisnewpayorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosonsitetradeisnewpayorderAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.onsite.trade.isnewpayorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosonsitetradeisnewpayorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosonsitetradeisnewpayorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutTradeNo is OutTradeNo Setter
// 外部订单号
func (r *AlibabamosonsitetradeisnewpayorderAPIRequest) SetOutTradeNo(_outTradeNo string) error {
	r._outTradeNo = _outTradeNo
	r.Set("out_trade_no", _outTradeNo)
	return nil
}

// GetOutTradeNo OutTradeNo Getter
func (r AlibabamosonsitetradeisnewpayorderAPIRequest) GetOutTradeNo() string {
	return r._outTradeNo
}
