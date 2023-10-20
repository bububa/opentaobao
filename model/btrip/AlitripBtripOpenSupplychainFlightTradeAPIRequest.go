package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopensupplychainflighttradeAPIRequest 【商旅】机票交易流水查询接口 API请求
// alitrip.btrip.open.supplychain.flight.trade
//
// 【商旅】杭州市政府机票交易流水接口查询
type AlitripbtripopensupplychainflighttradeAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenApiZzdSearchRq
}

// NewAlitripbtripopensupplychainflighttradeRequest 初始化AlitripbtripopensupplychainflighttradeAPIRequest对象
func NewAlitripbtripopensupplychainflighttradeRequest() *AlitripbtripopensupplychainflighttradeAPIRequest {
	return &AlitripbtripopensupplychainflighttradeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopensupplychainflighttradeAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.supplychain.flight.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopensupplychainflighttradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopensupplychainflighttradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripopensupplychainflighttradeAPIRequest) SetRq(_rq *OpenApiZzdSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopensupplychainflighttradeAPIRequest) GetRq() *OpenApiZzdSearchRq {
	return r._rq
}
