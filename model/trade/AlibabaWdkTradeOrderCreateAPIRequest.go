package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeOrderCreateAPIRequest 外部交易订单创单接口 API请求
// alibaba.wdk.trade.order.create
//
// 通过该接口可以再盒马创建交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
type AlibabaWdkTradeOrderCreateAPIRequest struct {
	model.Params
	// 待创建的订单
	_trade *TradeOrder
}

// NewAlibabaWdkTradeOrderCreateRequest 初始化AlibabaWdkTradeOrderCreateAPIRequest对象
func NewAlibabaWdkTradeOrderCreateRequest() *AlibabaWdkTradeOrderCreateAPIRequest {
	return &AlibabaWdkTradeOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeOrderCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Trade Setter
// 待创建的订单
func (r *AlibabaWdkTradeOrderCreateAPIRequest) SetTrade(_trade *TradeOrder) error {
	r._trade = _trade
	r.Set("trade", _trade)
	return nil
}

// Get Trade Getter
func (r AlibabaWdkTradeOrderCreateAPIRequest) GetTrade() *TradeOrder {
	return r._trade
}
