package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopencostcentermodifyAPIRequest 修改成本中心 API请求
// alitrip.btrip.open.cost.center.modify
//
// 修改成本中心
type AlitripbtripopencostcentermodifyAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterModifyRq
}

// NewAlitripbtripopencostcentermodifyRequest 初始化AlitripbtripopencostcentermodifyAPIRequest对象
func NewAlitripbtripopencostcentermodifyRequest() *AlitripbtripopencostcentermodifyAPIRequest {
	return &AlitripbtripopencostcentermodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopencostcentermodifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopencostcentermodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopencostcentermodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripopencostcentermodifyAPIRequest) SetRq(_rq *OpenCostCenterModifyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopencostcentermodifyAPIRequest) GetRq() *OpenCostCenterModifyRq {
	return r._rq
}
