package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcostcentermodifyAPIRequest 修改外部成本中心 API请求
// alitrip.btrip.cost.center.modify
//
// 修改外部成本中心，设置成员，设置支付宝账号，设置名称，编号等
type AlitripbtripcostcentermodifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterModifyRq
}

// NewAlitripbtripcostcentermodifyRequest 初始化AlitripbtripcostcentermodifyAPIRequest对象
func NewAlitripbtripcostcentermodifyRequest() *AlitripbtripcostcentermodifyAPIRequest {
	return &AlitripbtripcostcentermodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcostcentermodifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcostcentermodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcostcentermodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripcostcentermodifyAPIRequest) SetRq(_rq *OpenCostCenterModifyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcostcentermodifyAPIRequest) GetRq() *OpenCostCenterModifyRq {
	return r._rq
}
