package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktradeorderqueryAPIRequest 查询外部交易订单接口 API请求
// alibaba.wdk.trade.order.query
//
// 通过该接口可以在盒马查询交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
type AlibabawdktradeorderqueryAPIRequest struct {
	model.Params
	// 订单查询
	_query *TradeOrderQuery
}

// NewAlibabawdktradeorderqueryRequest 初始化AlibabawdktradeorderqueryAPIRequest对象
func NewAlibabawdktradeorderqueryRequest() *AlibabawdktradeorderqueryAPIRequest {
	return &AlibabawdktradeorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdktradeorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trade.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdktradeorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdktradeorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 订单查询
func (r *AlibabawdktradeorderqueryAPIRequest) SetQuery(_query *TradeOrderQuery) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r AlibabawdktradeorderqueryAPIRequest) GetQuery() *TradeOrderQuery {
	return r._query
}
