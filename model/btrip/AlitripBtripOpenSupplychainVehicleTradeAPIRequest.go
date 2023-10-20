package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopensupplychainvehicletradeAPIRequest 商旅用车交易流水接口 API请求
// alitrip.btrip.open.supplychain.vehicle.trade
//
// 商旅用车交易流水接口
type AlitripbtripopensupplychainvehicletradeAPIRequest struct {
	model.Params
	// 入参
	_rq *OpenApiZzdSearchRq
}

// NewAlitripbtripopensupplychainvehicletradeRequest 初始化AlitripbtripopensupplychainvehicletradeAPIRequest对象
func NewAlitripbtripopensupplychainvehicletradeRequest() *AlitripbtripopensupplychainvehicletradeAPIRequest {
	return &AlitripbtripopensupplychainvehicletradeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopensupplychainvehicletradeAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.supplychain.vehicle.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopensupplychainvehicletradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopensupplychainvehicletradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参
func (r *AlitripbtripopensupplychainvehicletradeAPIRequest) SetRq(_rq *OpenApiZzdSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopensupplychainvehicletradeAPIRequest) GetRq() *OpenApiZzdSearchRq {
	return r._rq
}
