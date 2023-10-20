package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcostcenternewAPIRequest 新建外部成本中心 API请求
// alitrip.btrip.cost.center.new
//
// 新建外部成本中心
type AlitripbtripcostcenternewAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterSaveRq
}

// NewAlitripbtripcostcenternewRequest 初始化AlitripbtripcostcenternewAPIRequest对象
func NewAlitripbtripcostcenternewRequest() *AlitripbtripcostcenternewAPIRequest {
	return &AlitripbtripcostcenternewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcostcenternewAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcostcenternewAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcostcenternewAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripcostcenternewAPIRequest) SetRq(_rq *OpenCostCenterSaveRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcostcenternewAPIRequest) GetRq() *OpenCostCenterSaveRq {
	return r._rq
}
