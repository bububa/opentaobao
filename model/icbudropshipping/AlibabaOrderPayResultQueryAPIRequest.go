package icbudropshipping

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaorderpayresultqueryAPIRequest alibaba查询订单支付结果 API请求
// alibaba.order.pay.result.query
//
// alibaba查询订单支付结果
type AlibabaorderpayresultqueryAPIRequest struct {
	model.Params
	// order id
	_tradeId int64
}

// NewAlibabaorderpayresultqueryRequest 初始化AlibabaorderpayresultqueryAPIRequest对象
func NewAlibabaorderpayresultqueryRequest() *AlibabaorderpayresultqueryAPIRequest {
	return &AlibabaorderpayresultqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaorderpayresultqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.order.pay.result.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaorderpayresultqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaorderpayresultqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTradeId is TradeId Setter
// order id
func (r *AlibabaorderpayresultqueryAPIRequest) SetTradeId(_tradeId int64) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// GetTradeId TradeId Getter
func (r AlibabaorderpayresultqueryAPIRequest) GetTradeId() int64 {
	return r._tradeId
}
