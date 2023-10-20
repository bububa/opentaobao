package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcostcenterentityaddAPIRequest 增加外部成本中心人员信息 API请求
// alitrip.btrip.cost.center.entity.add
//
// 增加外部成本中心人员信息
type AlitripbtripcostcenterentityaddAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterAddEntityRq
}

// NewAlitripbtripcostcenterentityaddRequest 初始化AlitripbtripcostcenterentityaddAPIRequest对象
func NewAlitripbtripcostcenterentityaddRequest() *AlitripbtripcostcenterentityaddAPIRequest {
	return &AlitripbtripcostcenterentityaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcostcenterentityaddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.entity.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcostcenterentityaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcostcenterentityaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripcostcenterentityaddAPIRequest) SetRq(_rq *OpenCostCenterAddEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcostcenterentityaddAPIRequest) GetRq() *OpenCostCenterAddEntityRq {
	return r._rq
}
