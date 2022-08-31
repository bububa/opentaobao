package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripFlightOrderSearchAPIRequest 获取机票订单列表 API请求
// alitrip.btrip.flight.order.search
//
// 第三方OA厂商获取机票订单列表
type AlitripBtripFlightOrderSearchAPIRequest struct {
	model.Params
	// 请求
	_rq *OpenSearchRq
}

// NewAlitripBtripFlightOrderSearchRequest 初始化AlitripBtripFlightOrderSearchAPIRequest对象
func NewAlitripBtripFlightOrderSearchRequest() *AlitripBtripFlightOrderSearchAPIRequest {
	return &AlitripBtripFlightOrderSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripFlightOrderSearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripFlightOrderSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRq is Rq Setter
// 请求
func (r *AlitripBtripFlightOrderSearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripFlightOrderSearchAPIRequest) GetRq() *OpenSearchRq {
	return r._rq
}
