package trade

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkTradeOrderCreateAPIRequest) Reset() {
	r._trade = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkTradeOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkTradeOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkTradeOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrade is Trade Setter
// 待创建的订单
func (r *AlibabaWdkTradeOrderCreateAPIRequest) SetTrade(_trade *TradeOrder) error {
	r._trade = _trade
	r.Set("trade", _trade)
	return nil
}

// GetTrade Trade Getter
func (r AlibabaWdkTradeOrderCreateAPIRequest) GetTrade() *TradeOrder {
	return r._trade
}

var poolAlibabaWdkTradeOrderCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkTradeOrderCreateRequest()
	},
}

// GetAlibabaWdkTradeOrderCreateRequest 从 sync.Pool 获取 AlibabaWdkTradeOrderCreateAPIRequest
func GetAlibabaWdkTradeOrderCreateAPIRequest() *AlibabaWdkTradeOrderCreateAPIRequest {
	return poolAlibabaWdkTradeOrderCreateAPIRequest.Get().(*AlibabaWdkTradeOrderCreateAPIRequest)
}

// ReleaseAlibabaWdkTradeOrderCreateAPIRequest 将 AlibabaWdkTradeOrderCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkTradeOrderCreateAPIRequest(v *AlibabaWdkTradeOrderCreateAPIRequest) {
	v.Reset()
	poolAlibabaWdkTradeOrderCreateAPIRequest.Put(v)
}
