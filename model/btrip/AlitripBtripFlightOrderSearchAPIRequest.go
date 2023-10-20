package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripflightordersearchAPIRequest 获取机票订单列表 API请求
// alitrip.btrip.flight.order.search
//
// 第三方OA厂商获取机票订单列表
type AlitripbtripflightordersearchAPIRequest struct {
	model.Params
	// 请求
	_rq *OpenSearchRq
}

// NewAlitripbtripflightordersearchRequest 初始化AlitripbtripflightordersearchAPIRequest对象
func NewAlitripbtripflightordersearchRequest() *AlitripbtripflightordersearchAPIRequest {
	return &AlitripbtripflightordersearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripflightordersearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.flight.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripflightordersearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripflightordersearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求
func (r *AlitripbtripflightordersearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripflightordersearchAPIRequest) GetRq() *OpenSearchRq {
	return r._rq
}
