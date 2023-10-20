package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopencostcenterentityaddAPIRequest 增加成本中心人员信息 API请求
// alitrip.btrip.open.cost.center.entity.add
//
// 增加成本中心人员信息
type AlitripbtripopencostcenterentityaddAPIRequest struct {
	model.Params
	// 入参对象
	_rq *OpenCostCenterAddEntityRq
}

// NewAlitripbtripopencostcenterentityaddRequest 初始化AlitripbtripopencostcenterentityaddAPIRequest对象
func NewAlitripbtripopencostcenterentityaddRequest() *AlitripbtripopencostcenterentityaddAPIRequest {
	return &AlitripbtripopencostcenterentityaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripopencostcenterentityaddAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.open.cost.center.entity.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripopencostcenterentityaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripopencostcenterentityaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRq is Rq Setter
// 入参对象
func (r *AlitripbtripopencostcenterentityaddAPIRequest) SetRq(_rq *OpenCostCenterAddEntityRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// GetRq Rq Getter
func (r AlitripbtripopencostcenterentityaddAPIRequest) GetRq() *OpenCostCenterAddEntityRq {
	return r._rq
}
