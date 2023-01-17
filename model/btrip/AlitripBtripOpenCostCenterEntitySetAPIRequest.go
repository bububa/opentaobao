package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenCostCenterEntitySetAPIRequest 设置成本中心人员信息 API请求
// alitrip.btrip.open.cost.center.entity.set
//
// 设置成本中心人员信息
type AlitripBtripOpenCostCenterEntitySetAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterSetEntityRq
}

// NewAlitripBtripOpenCostCenterEntitySetRequest 初始化AlitripBtripOpenCostCenterEntitySetAPIRequest对象
func NewAlitripBtripOpenCostCenterEntitySetRequest() *AlitripBtripOpenCostCenterEntitySetAPIRequest {
	return &AlitripBtripOpenCostCenterEntitySetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripOpenCostCenterEntitySetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.entity.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripOpenCostCenterEntitySetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripBtripOpenCostCenterEntitySetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripBtripOpenCostCenterEntitySetAPIRequest) SetRq(_rq *OpenCostCenterSetEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripBtripOpenCostCenterEntitySetAPIRequest) GetRq() *OpenCostCenterSetEntityRq {
	return r._rq
}
