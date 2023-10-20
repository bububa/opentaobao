package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcostcenterentitydeleteAPIRequest 删除外部成本中心人员信息 API请求
// alitrip.btrip.cost.center.entity.delete
//
// 删除外部成本中心人员信息
type AlitripbtripcostcenterentitydeleteAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterDeleteEntityRq
}

// NewAlitripbtripcostcenterentitydeleteRequest 初始化AlitripbtripcostcenterentitydeleteAPIRequest对象
func NewAlitripbtripcostcenterentitydeleteRequest() *AlitripbtripcostcenterentitydeleteAPIRequest {
	return &AlitripbtripcostcenterentitydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcostcenterentitydeleteAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.entity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcostcenterentitydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcostcenterentitydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripcostcenterentitydeleteAPIRequest) SetRq(_rq *OpenCostCenterDeleteEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcostcenterentitydeleteAPIRequest) GetRq() *OpenCostCenterDeleteEntityRq {
	return r._rq
}
