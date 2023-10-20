package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripRailTradeRefundAPIRequest 退票接口 API请求
// alitrip.rail.trade.refund
//
// 退票接口
type AlitripRailTradeRefundAPIRequest struct {
	model.Params
	// 入参
	_refundParam *RefundReq
}

// NewAlitripRailTradeRefundRequest 初始化AlitripRailTradeRefundAPIRequest对象
func NewAlitripRailTradeRefundRequest() *AlitripRailTradeRefundAPIRequest {
	return &AlitripRailTradeRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripRailTradeRefundAPIRequest) GetApiMethodName() string {
	return "alitrip.rail.trade.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripRailTradeRefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripRailTradeRefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefundParam is RefundParam Setter
// 入参
func (r *AlitripRailTradeRefundAPIRequest) SetRefundParam(_refundParam *RefundReq) error {
	r._refundParam = _refundParam
	r.Set("refund_param", _refundParam)
	return nil
}

// GetRefundParam RefundParam Getter
func (r AlitripRailTradeRefundAPIRequest) GetRefundParam() *RefundReq {
	return r._refundParam
}
