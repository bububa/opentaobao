package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripvehicleordersearchAPIRequest 用车订单查询接口 API请求
// alitrip.btrip.vehicle.order.search
//
// 企业获取商旅用车订单数据
type AlitripbtripvehicleordersearchAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenSearchRq
}

// NewAlitripbtripvehicleordersearchRequest 初始化AlitripbtripvehicleordersearchAPIRequest对象
func NewAlitripbtripvehicleordersearchRequest() *AlitripbtripvehicleordersearchAPIRequest {
	return &AlitripbtripvehicleordersearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripvehicleordersearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.vehicle.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripvehicleordersearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripvehicleordersearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripvehicleordersearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripvehicleordersearchAPIRequest) GetRq() *OpenSearchRq {
	return r._rq
}
