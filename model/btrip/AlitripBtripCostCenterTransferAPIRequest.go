package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcostcentertransferAPIRequest 商旅成本中心转换为外部成本中心 API请求
// alitrip.btrip.cost.center.transfer
//
// 商旅成本中心转换为外部成本中心
type AlitripbtripcostcentertransferAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterTransferRq
}

// NewAlitripbtripcostcentertransferRequest 初始化AlitripbtripcostcentertransferAPIRequest对象
func NewAlitripbtripcostcentertransferRequest() *AlitripbtripcostcentertransferAPIRequest {
	return &AlitripbtripcostcentertransferAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcostcentertransferAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.transfer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcostcentertransferAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcostcentertransferAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 请求对象
func (r *AlitripbtripcostcentertransferAPIRequest) SetRq(_rq *OpenCostCenterTransferRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripcostcentertransferAPIRequest) GetRq() *OpenCostCenterTransferRq {
	return r._rq
}
