package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcostcenterentitysetAPIRequest 设置外部成本中心人员信息 API请求
// alitrip.btrip.cost.center.entity.set
//
// 设置外部成本中心人员信息
type AlitripbtripcostcenterentitysetAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterSetEntityRq
}

// NewAlitripbtripcostcenterentitysetRequest 初始化AlitripbtripcostcenterentitysetAPIRequest对象
func NewAlitripbtripcostcenterentitysetRequest() *AlitripbtripcostcenterentitysetAPIRequest {
	return &AlitripbtripcostcenterentitysetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcostcenterentitysetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.entity.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcostcenterentitysetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcostcenterentitysetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripcostcenterentitysetAPIRequest) SetRq(_rq *OpenCostCenterSetEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcostcenterentitysetAPIRequest) GetRq() *OpenCostCenterSetEntityRq {
	return r._rq
}
