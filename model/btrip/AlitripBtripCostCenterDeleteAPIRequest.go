package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcostcenterdeleteAPIRequest 删除外部成本中心 API请求
// alitrip.btrip.cost.center.delete
//
// 删除外部成本中心
type AlitripbtripcostcenterdeleteAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterDeleteRq
}

// NewAlitripbtripcostcenterdeleteRequest 初始化AlitripbtripcostcenterdeleteAPIRequest对象
func NewAlitripbtripcostcenterdeleteRequest() *AlitripbtripcostcenterdeleteAPIRequest {
	return &AlitripbtripcostcenterdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcostcenterdeleteAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcostcenterdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcostcenterdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripcostcenterdeleteAPIRequest) SetRq(_rq *OpenCostCenterDeleteRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcostcenterdeleteAPIRequest) GetRq() *OpenCostCenterDeleteRq {
	return r._rq
}
