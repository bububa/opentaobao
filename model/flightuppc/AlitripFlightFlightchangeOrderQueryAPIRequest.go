package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripflightflightchangeorderqueryAPIRequest 订单维度航变查询 API请求
// alitrip.flight.flightchange.order.query
//
// 订单维度航变查询
type AlitripflightflightchangeorderqueryAPIRequest struct {
	model.Params
	// 航变信息查询请求体
	_queryFlightChangeOrderReq *QueryFlightChangeOrderReq
}

// NewAlitripflightflightchangeorderqueryRequest 初始化AlitripflightflightchangeorderqueryAPIRequest对象
func NewAlitripflightflightchangeorderqueryRequest() *AlitripflightflightchangeorderqueryAPIRequest {
	return &AlitripflightflightchangeorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripflightflightchangeorderqueryAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.flightchange.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripflightflightchangeorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripflightflightchangeorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryFlightChangeOrderReq is QueryFlightChangeOrderReq Setter
// 航变信息查询请求体
func (r *AlitripflightflightchangeorderqueryAPIRequest) SetQueryFlightChangeOrderReq(_queryFlightChangeOrderReq *QueryFlightChangeOrderReq) error {
	r._queryFlightChangeOrderReq = _queryFlightChangeOrderReq
	r.Set("query_flight_change_order_req", _queryFlightChangeOrderReq)
	return nil
}

// GetQueryFlightChangeOrderReq QueryFlightChangeOrderReq Getter
func (r AlitripflightflightchangeorderqueryAPIRequest) GetQueryFlightChangeOrderReq() *QueryFlightChangeOrderReq {
	return r._queryFlightChangeOrderReq
}
