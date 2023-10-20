package flightuppc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripFlightFlightchangeOrderQueryAPIRequest 订单维度航变查询 API请求
// alitrip.flight.flightchange.order.query
//
// 订单维度航变查询
type AlitripFlightFlightchangeOrderQueryAPIRequest struct {
	model.Params
	// 航变信息查询请求体
	_queryFlightChangeOrderReq *QueryFlightChangeOrderReq
}

// NewAlitripFlightFlightchangeOrderQueryRequest 初始化AlitripFlightFlightchangeOrderQueryAPIRequest对象
func NewAlitripFlightFlightchangeOrderQueryRequest() *AlitripFlightFlightchangeOrderQueryAPIRequest {
	return &AlitripFlightFlightchangeOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripFlightFlightchangeOrderQueryAPIRequest) GetApiMethodName() string {
	return "alitrip.flight.flightchange.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripFlightFlightchangeOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripFlightFlightchangeOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryFlightChangeOrderReq is QueryFlightChangeOrderReq Setter
// 航变信息查询请求体
func (r *AlitripFlightFlightchangeOrderQueryAPIRequest) SetQueryFlightChangeOrderReq(_queryFlightChangeOrderReq *QueryFlightChangeOrderReq) error {
	r._queryFlightChangeOrderReq = _queryFlightChangeOrderReq
	r.Set("query_flight_change_order_req", _queryFlightChangeOrderReq)
	return nil
}

// GetQueryFlightChangeOrderReq QueryFlightChangeOrderReq Getter
func (r AlitripFlightFlightchangeOrderQueryAPIRequest) GetQueryFlightChangeOrderReq() *QueryFlightChangeOrderReq {
	return r._queryFlightChangeOrderReq
}
