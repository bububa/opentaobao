package trade

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkTradeOrderCancelAPIRequest) Reset() {
	r._trade = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeOrderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkTradeOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrade is Trade Setter
// 待取消的订单
func (r *AlibabaWdkTradeOrderCancelAPIRequest) SetTrade(_trade *TradeOrder) error {
	r._trade = _trade
	r.Set("trade", _trade)
	return nil
}

// GetTrade Trade Getter
func (r AlibabaWdkTradeOrderCancelAPIRequest) GetTrade() *TradeOrder {
	return r._trade
}

var poolAlibabaWdkTradeOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkTradeOrderCancelRequest()
	},
}

// GetAlibabaWdkTradeOrderCancelRequest 从 sync.Pool 获取 AlibabaWdkTradeOrderCancelAPIRequest
func GetAlibabaWdkTradeOrderCancelAPIRequest() *AlibabaWdkTradeOrderCancelAPIRequest {
	return poolAlibabaWdkTradeOrderCancelAPIRequest.Get().(*AlibabaWdkTradeOrderCancelAPIRequest)
}

// ReleaseAlibabaWdkTradeOrderCancelAPIRequest 将 AlibabaWdkTradeOrderCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkTradeOrderCancelAPIRequest(v *AlibabaWdkTradeOrderCancelAPIRequest) {
	v.Reset()
	poolAlibabaWdkTradeOrderCancelAPIRequest.Put(v)
}
