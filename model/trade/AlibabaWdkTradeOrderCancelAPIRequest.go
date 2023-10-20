package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktradeordercancelAPIRequest 外部交易订单取消接口 API请求
// alibaba.wdk.trade.order.cancel
//
// 通过该接口可以再盒马取消交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
type AlibabawdktradeordercancelAPIRequest struct {
	model.Params
	// 待取消的订单
	_trade *TradeOrder
}

// NewAlibabawdktradeordercancelRequest 初始化AlibabawdktradeordercancelAPIRequest对象
func NewAlibabawdktradeordercancelRequest() *AlibabawdktradeordercancelAPIRequest {
	return &AlibabawdktradeordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdktradeordercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdktradeordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdktradeordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrade is Trade Setter
// 待取消的订单
func (r *AlibabawdktradeordercancelAPIRequest) SetTrade(_trade *TradeOrder) error {
	r._trade = _trade
	r.Set("trade", _trade)
	return nil
}

// GetTrade Trade Getter
func (r AlibabawdktradeordercancelAPIRequest) GetTrade() *TradeOrder {
	return r._trade
}
