package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeOrderCancelAPIRequest 外部交易订单取消接口 API请求
// alibaba.wdk.trade.order.cancel
//
// 通过该接口可以再盒马取消交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
type AlibabaWdkTradeOrderCancelAPIRequest struct {
	model.Params
	// 待取消的订单
	_trade *TradeOrder
}

// NewAlibabaWdkTradeOrderCancelRequest 初始化AlibabaWdkTradeOrderCancelAPIRequest对象
func NewAlibabaWdkTradeOrderCancelRequest() *AlibabaWdkTradeOrderCancelAPIRequest {
	return &AlibabaWdkTradeOrderCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeOrderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeOrderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Trade Setter
// 待取消的订单
func (r *AlibabaWdkTradeOrderCancelAPIRequest) SetTrade(_trade *TradeOrder) error {
	r._trade = _trade
	r.Set("trade", _trade)
	return nil
}

// Get Trade Getter
func (r AlibabaWdkTradeOrderCancelAPIRequest) GetTrade() *TradeOrder {
	return r._trade
}
