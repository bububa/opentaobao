package flightuppc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripFlightFlightchangeOrderQueryAPIRequest) Reset() {
	r._queryFlightChangeOrderReq = nil
	r.Params.ToZero()
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

var poolAlitripFlightFlightchangeOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripFlightFlightchangeOrderQueryRequest()
	},
}

// GetAlitripFlightFlightchangeOrderQueryRequest 从 sync.Pool 获取 AlitripFlightFlightchangeOrderQueryAPIRequest
func GetAlitripFlightFlightchangeOrderQueryAPIRequest() *AlitripFlightFlightchangeOrderQueryAPIRequest {
	return poolAlitripFlightFlightchangeOrderQueryAPIRequest.Get().(*AlitripFlightFlightchangeOrderQueryAPIRequest)
}

// ReleaseAlitripFlightFlightchangeOrderQueryAPIRequest 将 AlitripFlightFlightchangeOrderQueryAPIRequest 放入 sync.Pool
func ReleaseAlitripFlightFlightchangeOrderQueryAPIRequest(v *AlitripFlightFlightchangeOrderQueryAPIRequest) {
	v.Reset()
	poolAlitripFlightFlightchangeOrderQueryAPIRequest.Put(v)
}
