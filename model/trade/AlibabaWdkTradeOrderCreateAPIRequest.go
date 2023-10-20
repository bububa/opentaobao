package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktradeordercreateAPIRequest 外部交易订单创单接口 API请求
// alibaba.wdk.trade.order.create
//
// 通过该接口可以再盒马创建交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
type AlibabawdktradeordercreateAPIRequest struct {
	model.Params
	// 待创建的订单
	_trade *TradeOrder
}

// NewAlibabawdktradeordercreateRequest 初始化AlibabawdktradeordercreateAPIRequest对象
func NewAlibabawdktradeordercreateRequest() *AlibabawdktradeordercreateAPIRequest {
	return &AlibabawdktradeordercreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdktradeordercreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdktradeordercreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdktradeordercreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrade is Trade Setter
// 待创建的订单
func (r *AlibabawdktradeordercreateAPIRequest) SetTrade(_trade *TradeOrder) error {
	r._trade = _trade
	r.Set("trade", _trade)
	return nil
}

// GetTrade Trade Getter
func (r AlibabawdktradeordercreateAPIRequest) GetTrade() *TradeOrder {
	return r._trade
}
